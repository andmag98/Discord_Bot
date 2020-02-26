package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//WeatherHandler handles weather requests and returns a forecast
func WeatherHandler(city, countryCode string) (string, error) {

	client := http.DefaultClient

	url := "https://api.weatherbit.io/v2.0/current?city=" + city + "&country=" + countryCode + Key

	resp, err := DoRequest(url, client)
	if err != nil {
		return "```Could not get contact with API```", err
	}

	data := WeatherData{}

	//decodes data form the response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "```Could not find " + city + ". See !commands for help.```", err

	}

	defer resp.Body.Close()

	// convert to lower/upper case
	countryCode = strings.ToLower(countryCode)
	countryCodeOutput := strings.ToUpper(countryCode)

	// convert to fahrenheit
	appTempFahrenheit := ConvertToFahrenheit(data.Result[0].AppTemp)
	fahrenheit := ConvertToFahrenheit(data.Result[0].Temp)

	// convert to strings
	temp := fmt.Sprintf("%.f", data.Result[0].Temp)
	wind := fmt.Sprintf("%.2f", data.Result[0].Wind)
	appTempCelsius := fmt.Sprintf("%.1f", data.Result[0].AppTemp)

	//string output
	weather := ":flag_" + countryCode + ": | " + "**Weather for** " + city + ", " + countryCodeOutput + "\n" +
		"**Weather:** " + data.Result[0].WeatherStatus.Description + "\n" +
		":sunrise: **Sunrise: **" + data.Result[0].Sunrise + " UTC | :city_sunset: **Sunset: **" + data.Result[0].Sunset + " UTC\n" +
		"**Temp: **" + temp + "°C / " + fahrenheit + "°F | **Feels like: **" + appTempCelsius + "°C / " + appTempFahrenheit + "°F \n" +
		"**Wind: **" + data.Result[0].WindDir + " | " + "**Speed: **" + wind + " m/s\n"

	return weather, nil
}
