package src

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

// DoRequest handles all request and returns the response
func DoRequest(url string, c *http.Client) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Error in doRequest()")
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Error in doRequest()")
	}
	return resp, nil
}

//ConvertToFahrenheit converts celcius to fahrenheit
func ConvertToFahrenheit(celcius float32) string {
	far := (celcius * 1.8) + 32
	fahrenheit := fmt.Sprintf("%.1f", far)
	return fahrenheit
}

//Percent returns the percent the price represents relative to the sum
func Percent(sum, price int) float64 {
	percent := (float64(price) * 100) / float64(sum)
	return percent
}

//SetStatusEmoji Sets the status emoji for the diag info
func SetStatusEmoji(url string, c *http.Client) (emoji string, code int, err error) {
	resp, err := DoRequest(url, c)
	if err != nil {
		fmt.Println(err)
		return "", http.StatusInternalServerError, errors.Wrap(err, "Error in setStatusEmoji()")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ":x:", resp.StatusCode, nil
	}
	return ":white_check_mark:", resp.StatusCode, nil
}

//SendMessage sends a message passed as a string to Discord
func SendMessage(s *discordgo.Session, channelID, message string) error {

	_, err := s.ChannelMessageSend(channelID, message)
	if err != nil {
		return errors.Wrap(err, "Failed sanding message in SendMessage()")
	}
	return nil
}

//SendEmbedImageLink sends a message passed as a string to Discord
func SendEmbedImageLink(s *discordgo.Session, m *discordgo.MessageCreate, image string) error {

	embed := NewEmbed().SetImage(image)
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
	if err != nil {
		return errors.Wrap(err, "Failed sanding message in SendEmbedImageLink()")
	}
	return nil
}
