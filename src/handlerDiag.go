package src

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

//Diag displays diagnostics
func Diag() (message string, err error) {
	diagnostic := Diagnostic{}
	diagnostic.Uptime = int((time.Since(StartTime)).Seconds()) //uptime is calculated

	client := http.DefaultClient
	emoji, code, err := SetStatusEmoji("https://randomfox.ca/floof/", client) //Makes test query to fox API
	if err != nil {
		return "", errors.Wrap(err, "Error in Diag()")
	}
	diagnostic.FoxEmoji = emoji //Sets emoji and status code
	diagnostic.FoxCode = code

	emoji, code, err = SetStatusEmoji("https://api.weatherbit.io/v2.0/current?city=oslo&country=no"+Key, client) //Makes test query to weather API
	if err != nil {
		return "", errors.Wrap(err, "Error in Diag()")
	}
	diagnostic.WeatherEmoji = emoji //Sets emoji and status code
	diagnostic.WetherCode = code

	emoji, _, err = SetStatusEmoji("https://console.firebase.google.com/u/0/project/projectcloud-49e3d/database/firestore/data~2F", client) //Checks availability of database
	if err != nil {
		return "", errors.Wrap(err, "Error in Diag()")
	}
	diagnostic.DBEmoji = emoji //Sets emoji

	uptime := fmt.Sprintf("%d", diagnostic.Uptime) //Removes unneeded decimals from uptime

	return ("**Diagnostics**\n\n" +
		":heart: **Uptime: **" + uptime +
		" sec\n:fox: **Foxes API:**\t" + diagnostic.FoxEmoji +
		"\n:partly_sunny: **Weather API:**\t" + diagnostic.WeatherEmoji +
		"\n:file_cabinet: **Database:**\t" + diagnostic.DBEmoji), nil
}
