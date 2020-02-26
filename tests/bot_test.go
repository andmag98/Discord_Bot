package tests

import (
	"fmt"
	"project/src"
	"testing"

	"github.com/bwmarrin/discordgo"
)

//TestBot tests every possible combination of commands
func TestBot(t *testing.T) {
	src.ReadConfig("../config.json")
	err := src.InitDatabase("../")
	if err != nil {
		t.Errorf("Could init test database!")
	}
	err = src.ReminderInit("../")
	if err != nil {
		t.Errorf("Could init test reminder database!")
	}
	src.Start()

	channelID := "644859571780976660" //Id to testing channel in discord where the output will end up

	fmt.Println(src.Token)
	goBot, err := discordgo.New("Bot " + src.Token)
	if err != nil {
		fmt.Println(err.Error())
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	src.BotID = u.ID

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer goBot.Close()

	fmt.Println("Bot is running!")

	_, err = goBot.ChannelMessageSend(channelID, "!bills")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!bills snerk")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!bills register 100 other")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!bills sum")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!bills sum other")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!bills sum jsh sdjhfb dsjhb sdyhjb")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!bills diagram")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!bills diagram fggh")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!bills total")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!bills total hgf")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!fox")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!weather oslo no")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!weather")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!help")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!help weather")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!help remind")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!help bills")
	if err != nil {
		t.Error(err)
	}

	_, err = goBot.ChannelMessageSend(channelID, "!remind pull")
	if err != nil {
		t.Error(err)
	}

	_, err = goBot.ChannelMessageSend(channelID, "!remind 05/00/00 hei")
	if err != nil {
		t.Error(err)
	}

	_, err = goBot.ChannelMessageSend(channelID, "!remind 5/00 hei")
	if err != nil {
		t.Error(err)
	}

	_, err = goBot.ChannelMessageSend(channelID, "!diag")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!snerk")
	if err != nil {
		t.Error(err)
	}
	_, err = goBot.ChannelMessageSend(channelID, "!weather langhus no")
	if err != nil {
		t.Error(err)
	}

	err = src.SendMessage(goBot, channelID, "hei")
	if err != nil {
		t.Error(err)
	}

}
