package src

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// BotID variable
var BotID string

// Start bot
func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bot is running!")

}

//messageHandler handles all base commands
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Message from user has "!":
	if strings.HasPrefix(m.Content, config.BotPrefix) {

		parts := strings.Fields(m.Content)
		command := parts[0]
		channelID := m.ChannelID

		switch command {
		case "!bills":
			if len(parts) == 1 {
				defaultMenu(s, m)
			} else if len(parts) != 1 {
				billCommands(s, m, parts)
			}
		case "!help":
			if len(parts) == 1 {
				defaultMenu(s, m)
			} else if len(parts) != 1 {
				helpCommands(s, m, parts)
			}
		case "!weather":
			if len(parts) != 3 {
				weatherMenu(s, m)
			} else {
				message, _ := WeatherHandler(parts[1], parts[2])
				err := SendMessage(s, channelID, message)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		case "!diag":
			msg, err := Diag()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = SendMessage(s, channelID, msg)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "!fox":
			img, err := RandomFox()
			if err == nil {
				err := SendEmbedImageLink(s, m, img)
				if err != nil {
					return
				}
			}
		case "!remind":
			reminderCommands(s, m, parts)
		default:
			defaultMenu(s, m)
		}
	}
}

//billCommands handles all commands specific to bills
func billCommands(s *discordgo.Session, m *discordgo.MessageCreate, parts []string) {

	command := parts[1]
	userID := m.Author.ID
	userName := m.Author.Username
	channelID := m.ChannelID

	switch command {
	case "register":
		if len(parts) != 4 {
			billMenu(s, m)
		} else {
			message, _ := RegisterBill(userID, userName, parts[2], parts[3])
			err := SendMessage(s, channelID, message)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "sum":
		if len(parts) == 2 {
			message := SumAllHandler(userID, userName)
			err := SendMessage(s, channelID, message)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else if len(parts) == 3 {
			message := SumTypeHandler(userID, userName, parts[2])
			err := SendMessage(s, channelID, message)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			billMenu(s, m)
		}
	case "diagram":
		if len(parts) == 2 {
			file, text := DiagramHandler(userID)
			_, _ = s.ChannelFileSend(m.ChannelID, text, file)
		} else {
			billMenu(s, m)
		}
	case "total":
		if len(parts) == 2 {
			message := TotalHandler(userID)
			err := SendMessage(s, channelID, message)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			billMenu(s, m)
		}
	default:
		billMenu(s, m)
	}
}

//helpCommands handles all commands specific to help
func helpCommands(s *discordgo.Session, m *discordgo.MessageCreate, parts []string) {

	command := parts[1]

	switch command {
	case "bills":
		billMenu(s, m)
	case "reminder":
		reminderMenu(s, m)
	case "weather":
		weatherMenu(s, m)
	default:
		defaultMenu(s, m)
	}
}

//reminderCommands handles all commands specific to reminder
func reminderCommands(s *discordgo.Session, m *discordgo.MessageCreate, parts []string) {
	if len(parts) == 2 && parts[1] == "pull" {
		err := restartReminders(s, m)
		if err != nil {
			_ = SendMessage(s, m.ChannelID, "Couldn't pull reminders from DB")
			fmt.Println(err)
		}
	} else {
		//If everything is fine this error will be nil:
		var err error
		//Checks for the following timestamp in the input from user:
		if len(parts) == 3 {
			err = reminderNewHandler(s, m, parts)
			if err != nil {
				_ = SendMessage(s, m.ChannelID, ":x: Couldn't make reminder")
			}
		} else {
			err := SendMessage(s, m.ChannelID, ":x: Something went wrong! write: !remind sec/min/hours message(with no spaces)")
			if err != nil {
				return
			}
		}
	}
}
