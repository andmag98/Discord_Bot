package src

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

//List of all reminders:
var reminders []Reminder = make([]Reminder, 0)

//RunningReminder is the struct for the running reminders:
type RunningReminder struct {
	//Message to remind
	message string
	//user ID:
	UserID string
	//time at which to remind
	timestampToRemind time.Time
	//time at which the struct was created (if you call NewReminder())
	//on creating it:
	timestampStart time.Time
	//Timer of duration  between start and remind:
	stopWatch *time.Timer
	channelID string
	dur       time.Duration
}

//NewRunningReminder is basically the constructor for runningReminder
func NewRunningReminder(channelID string, message string, userID string, discriminator string, seconds int, minutes int, hours int) (*RunningReminder, error) {
	var err error
	var rem *RunningReminder = nil
	//Checks if any of the time parameters are greater than 0:
	if hours >= 0 && minutes >= 0 && seconds >= 0 {
		if len(channelID) > 0 {
			var r RunningReminder
			//Set the timestamp:
			r.timestampStart = time.Now()
			//Creates a duration of the parameters to be reminded:
			d := time.Duration(time.Duration(hours) * time.Hour)
			d = d + time.Duration(time.Duration(minutes)*time.Minute)
			d = d + time.Duration(time.Duration(seconds)*time.Second)
			r.dur = d
			//Set the timestamp to be reminded
			r.timestampToRemind = r.timestampStart.Add(d)
			difference := r.timestampToRemind.Sub(r.timestampStart)
			//starts the timer with the duration from "difference":
			r.stopWatch = time.NewTimer(difference)
			//Set the message:
			if len(message) != 0 {
				r.message = message
			} else {
				r.message = "Didn't insert message"
			}
			r.channelID = channelID
			if userID != "" && discriminator != "" {
				r.UserID = userID
			} else {
				r.UserID = ""
			}

			rem = &r
		} else {
			err = errors.New("Invalid function parameters. channelID can't be empty")
		}
	} else {
		err = errors.New("Invalid function parameters. hours, minutes or seconds are below 0")
	}
	return rem, err
}

//NewRunningReminderDB this is used when the reminders are pulled from the DB:
func NewRunningReminderDB(message string, userID string, remindTime *time.Time) (*RunningReminder, error) {
	var err error
	var rem *RunningReminder = nil
	if len(message) > 0 && remindTime != nil {
		var r RunningReminder
		//Sets the timestamp
		r.message = message
		r.UserID = userID
		r.timestampStart = time.Now()
		r.timestampToRemind = *remindTime
		r.dur = r.timestampToRemind.Sub(r.timestampStart)

		//If the timestampToBeReminded is passed while the bot was turned off:
		if r.dur.Nanoseconds() > 0 {
			r.stopWatch = time.NewTimer(r.dur)
		} else {
			r.message = "NOTE! Couldn't send reminder at time: " + message
			r.stopWatch = time.NewTimer(2 * time.Second)
		}
		rem = &r
	} else {
		err = errors.New("Function called without all using both parameters")
	}
	return rem, err
}

//reminderNewHandler handles new reminders
func reminderNewHandler(s *discordgo.Session, m *discordgo.MessageCreate, parts []string) error {
	var err error
	//Splits the timestamp into Sec/Min/Hours
	myTimeStamp := strings.Split(parts[1], "/")
	if len(myTimeStamp) == 3 {
		//array to stringconvert the string array:
		var myTimeStampIntified []int
		//Goes through each timestamp:
		for i := range myTimeStamp {
			var t int //time
			//Convert input from user to int (if possible):
			t, err = strconv.Atoi(myTimeStamp[i])
			if err == nil {
				//Append the int to the array
				myTimeStampIntified = append(myTimeStampIntified, t)
			} else {
				//Break out of for loop:
				break
			}
		}

		//If all the inputs was valid:
		if err == nil {
			//fmt.Println(user)
			myReminder, err := NewRunningReminder(m.ChannelID, parts[2], m.Author.ID,
				m.Author.Discriminator, myTimeStampIntified[0], myTimeStampIntified[1], myTimeStampIntified[2])
			if err == nil {
				fmt.Println(myReminder.timestampToRemind.String())
				//Adds the reminder to the reminder DB:
				//Creates a new reminder to be added to the list:
				r, err := NewReminder(m.ChannelID, myReminder.message, myReminder.UserID, &myReminder.timestampToRemind)
				if err == nil {
					//Adds the new reminder to the list:
					err = AddReminderToDB(r)
					if err == nil {
						reminders = append(reminders, *r)
						//Adds the reminder to Firestore database:

						//DEBUG//
						fmt.Println("r content:")
						fmt.Println(r.Message)
						fmt.Println(r.Date)
						fmt.Println(r.ChannelID)
						//DEBUG//

						//Start running code on the side:
						go func() {

							err = runReminder(s, m, myReminder)
							//Deletes reminder from firebase:
							if err == nil {
								//DETELE//
								fmt.Println(r.DocID)
								//DELETE//

								//Deletes reminder r from list of reminders:
								err = r.DeleteReminderFromDB()
								if err == nil {
									deleteReminder(r)
								}
							}
						}()
					} else {
						_ = SendMessage(s, m.ChannelID, ":x: **Didn't manage to add reminder to DB!**")
					}
				}
			}
		}
	}
	return err
}

//restartReminders used for pulling down reminders from reminders list
func restartReminders(s *discordgo.Session, m *discordgo.MessageCreate) error {
	var err error = nil
	//Goes through all the reminders:
	count := 0
	for i := 0; i < len(reminders); i++ {
		fmt.Println(reminders[i].ChannelID)
		fmt.Println(m.ChannelID)
		//If it finds a reminder that is from the current channel AND has not been started
		//(AKA it was just pulled from DB after bot reboot):
		if reminders[i].PulledFromDB && m.ChannelID == reminders[i].ChannelID {
			//Converts the time string to a time type:
			d := strings.Split(reminders[i].Date, " m")[0]
			t, err := time.Parse(layout, d)
			//If there were no errors:
			if err == nil {
				//Sets a new running reminder:
				myReminder, err := NewRunningReminderDB(reminders[i].Message, reminders[i].UserID, &t)
				if err == nil {

					//DEBUG//
					fmt.Println(reminders[i].DocID)
					fmt.Println(myReminder.timestampToRemind.String())
					//DEBUG//

					reminders[i].PulledFromDB = false
					index := i
					//Start running code on the side:
					go func() {
						err = runReminder(s, m, myReminder)
						if err == nil {
							//Deletes reminder from firebase:
							err = reminders[index].DeleteReminderFromDB()
							if err == nil {
								//Deletes reminder r from list of reminders:
								reminders[index].toBeDeleted = true
							}
						}
					}()
					count++
				}
			} else {
				_ = SendMessage(s, m.ChannelID, ":x: **Couldn't parse time!**")
				fmt.Println(err)
			}
		}
	}
	deleteReminders()
	if count == 0 {
		_ = SendMessage(s, m.ChannelID, ":x: **No reminders were started!**")
	}
	return err
}

//runReminder rund reminder
func runReminder(s *discordgo.Session, m *discordgo.MessageCreate, r *RunningReminder) error {
	var err error
	if s != nil && m != nil && r != nil {

		var hrs, min, sec string
		if int(r.dur.Hours()) > 0 {
			hrs = strconv.Itoa(int(r.dur.Hours())) + "h "
		}
		if int(r.dur.Minutes()) > 0 {
			min = strconv.Itoa(int(r.dur.Minutes())%60) + "m "
		}
		if int(r.dur.Seconds()) > 0 {
			sec = strconv.Itoa(int(r.dur.Seconds())%60) + "s"
		}

		//INFORMING USER:
		err := SendMessage(s, m.ChannelID, ":white_check_mark: **Reminder for: **<@"+r.UserID+">\n"+
			":incoming_envelope: **Message**: "+r.message+"\n"+
			":timer: **Will go off in**: "+hrs+min+sec+"\n")
		if err != nil {
			return errors.Wrap(err, "Error in runReminder()")
		}
		//When the stopWatch timer is done:
		<-r.stopWatch.C

		err = SendMessage(s, m.ChannelID, ":bell: **Reminder for: **<@"+r.UserID+">\n"+
			":mailbox_with_mail: **Message**: "+r.message)
		if err != nil {
			return errors.Wrap(err, "Error in runReminder()")
		}
	} else {
		err = errors.New("Invalid Parameters")
	}
	return err
}
