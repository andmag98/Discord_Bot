package src

import (
	"github.com/bwmarrin/discordgo"
)

//defaultMenu prints the default menu
func defaultMenu(s *discordgo.Session, m *discordgo.MessageCreate) {

	help := space(20)
	diag := space(75)
	fox := space(78)

	helpMenu := "```css\n Command list\n```" + "\n" +
		"**:pray: \t !help** `bills` | `reminder` | `weather`" + help + " - !help + command \n" +
		"**:heart: \t !diag** " + diag + " - Uptime \n" +
		"**:fox: \t !fox** " + fox + "- Random pictures of foxes \n" +
		"\n"

	_, _ = s.ChannelMessageSend(m.ChannelID, helpMenu)
}

//billMenu prints bill menu
func billMenu(s *discordgo.Session, m *discordgo.MessageCreate) {

	syntax := space(45)
	register := space(8)
	sum := space(77)
	diagram := space(93)
	total := space(100)

	helpBills := "```css\n Bill command list\n```" + "\n" +
		"Syntax: ![command] `parameter1` `parameter2` " + syntax + "- <Description of command>\n\n" +
		"**:credit_card: \t !bills register -** `money` `food/clothes/electronics/other` " + register + " - Register a bill \n" +
		"**:dollar: \t !bills sum -**  OPT: `type`" + sum + " - Sums bills\n" +
		"**:bar_chart: \t !bills diagram** " + diagram + " - Purchaces in a PIE-chart\n" +
		"**:dividers: \t !bills total** " + total + " - Overview of total spending" +
		"\n"

	_, _ = s.ChannelMessageSend(m.ChannelID, helpBills)
}

//Reminder functionality instructions:
func reminderMenu(s *discordgo.Session, m *discordgo.MessageCreate) {

	syntax := space(27)
	remind := space(35)
	rstrtReminders := space(69)

	helpReminder := "```css\n Reminder command list\n```" + "\n" +
		"Syntax: ![command] `parameter1` `parameter2` " + syntax + "- <Description of command>\n\n" +
		"**:bell: \t !remind** `<sek/min/hrs>` `message` " + remind + "- Register a reminder \n" +
		"**:bell: \t !remind pull** " + rstrtReminders + "- Pulls reminders from DB after reboot\n"
	_, _ = s.ChannelMessageSend(m.ChannelID, helpReminder)
}

//weatherMenu prints weather menu
func weatherMenu(s *discordgo.Session, m *discordgo.MessageCreate) {

	syntax := space(30)
	weather := space(45)

	helpWeather := "```css\n Weather command list\n```" + "\n" +
		"Syntax: ![command] `parameter1` `parameter2` " + syntax + "- <Description of command>\n\n" +
		"**:white_sun_small_cloud: \t !weather -** `city` `country code` " + weather + " - Weather in a city \n"

	_, _ = s.ChannelMessageSend(m.ChannelID, helpWeather)
}

//space returns a string of spaces with the specified number of spaces
func space(number int) string {

	var noOfSpace string

	for i := 0; i < number; i++ {
		noOfSpace += " "
	}

	return noOfSpace
}
