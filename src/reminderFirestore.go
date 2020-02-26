package src

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const layout string = "2006-01-02 15:04:05.999999999 -0700 MST"

//RemindCtx Firebase context and client used by Firebase functions throughout the program.
var RemindCtx context.Context

//RemindClient client
var RemindClient *firestore.Client = nil

//NewReminder creates a reminder (that is to be added to reminder list):
func NewReminder(channelID string, m string, u string, remindDate *time.Time) (*Reminder, error) {
	var rem *Reminder = nil
	var err error
	if len(channelID) > 0 && len(m) > 0 && remindDate != nil {
		var r Reminder

		r.Message = m
		r.UserID = u
		r.Date = remindDate.String()
		r.ChannelID = channelID
		r.PulledFromDB = false
		r.toBeDeleted = false
		rem = &r
	} else {
		err = errors.New("Invalid function parameters. channelID, m and remindDate has to be filled")
	}
	return rem, err
}

//AddReminderToDB adds the reminder to the DB:
func AddReminderToDB(r *Reminder) error {

	myRef, _, err := RemindClient.Collection("reminders").Add(RemindCtx, map[string]interface{}{
		"Message":   r.Message,
		"Date":      r.Date,
		"ChannelID": r.ChannelID,
		"UserID":    r.UserID,
	})
	if err == nil {
		r.DocID = myRef.ID
	} else {
		fmt.Println(err)
	}
	return errors.Wrap(err, "AddReminderToDB() failed when RemindClient.Collection(...).Add(...)")
}

//DeleteReminderFromDB deletes the reminder from DB:
func (r *Reminder) DeleteReminderFromDB() error {
	_, err := RemindClient.Collection("reminders").Doc(r.DocID).Delete(RemindCtx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	return errors.Wrap(err, "DeleteReminderToDB() failed when RemindClient.Collection(...).Doc(...).Delete(...)")
}

//Delete all flagged reminders:
func deleteReminders() {
	for i := 0; i < len(reminders); i++ {
		//If it finds the reminder to be deleted:
		if reminders[i].toBeDeleted {
			//Removes the slot i (NOTE NOT STABLE)
			reminders[i] = reminders[len(reminders)-1]
			reminders = reminders[:len(reminders)-1]
			break
		}
	}
}

//Delete a single reminder:
func deleteReminder(r *Reminder) {
	for i := range reminders {
		//If it finds the reminder to be deleted:
		if reminders[i].DocID == r.DocID {
			//Removes the slot i (NOTE NOT STABLE)
			reminders[i] = reminders[len(reminders)-1]
			reminders = reminders[:len(reminders)-1]
			break
		}
	}
}

//GetReminderFromDB gets all the reminders from DB and sets it to the list of reminders:
func GetReminderFromDB() ([]Reminder, error) {
	remList := make([]Reminder, 0)
	var err error
	iter := RemindClient.Collection("reminders").Documents(RemindCtx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return remList, errors.Wrap(err, "GetReminerFromDB() something went wrong iter.Next()")
		}
		m := doc.Data()
		var r Reminder
		//ok used to check if casting goes well:
		var ok bool
		fmt.Println(m)
		r.DocID = doc.Ref.ID
		//Casts all the rows from map:
		r.Message, ok = m["Message"].(string)
		if ok {
			r.Date, ok = m["Date"].(string)
			if ok {
				r.ChannelID, ok = m["ChannelID"].(string)
				if ok {
					r.UserID, ok = m["UserID"].(string)
					if ok {
						r.PulledFromDB = true
						r.toBeDeleted = false
						remList = append(remList, r)
						//DEBUG//
						fmt.Println(r)
						//DEBUG//
					}
				}
			}
		}
	}
	return remList, err
}

//ReminderInit pulls all the requests from DB (NOTE! FC is used because we have moved _test to a seperate package):
func ReminderInit(filePath string) error {
	var err error
	// Firebase initialisation
	RemindCtx = context.Background()

	myServiceAccount := option.WithCredentialsFile(filePath + FirestoreCredential)
	app, err := firebase.NewApp(RemindCtx, nil, myServiceAccount)
	if err == nil {

		RemindClient, err = app.Firestore(RemindCtx)

		if err == nil {
			reminders, err = GetReminderFromDB()
			if err != nil {
				fmt.Println("GetReminderFromDB() got an error: ")
				fmt.Println(err)
			}
		}
	}
	if RemindClient != nil {
		//DELETE//
		fmt.Println("remindClient is running")
		//DELETE//
	}
	return err
}
