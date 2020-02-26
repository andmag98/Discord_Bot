package tests

import (
	"project/src"
	"testing"
)

func TestWeather(t *testing.T) {

	text, err := src.WeatherHandler("langhus", "no")
	if err == nil {
		t.Error(err)
	}
	if text != "```Could not find langhus. See !commands for help.```" {
		t.Error(err)
	}

	text2, err := src.WeatherHandler("oslo", "no")
	if err != nil {
		t.Error(err)
	}
	if text2 == "```Could not find oslo. See !commands for help.```" {
		t.Error(err)
	}

}
