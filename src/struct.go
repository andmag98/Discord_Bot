package src

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// Fox struct: to get random pictures of foxes
type Fox struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}

//Weather struct to collect weather data in a city
type Weather struct {
	Temp          float32 `json:"temp"`
	Sunset        string  `json:"sunset"`
	Sunrise       string  `json:"sunrise"`
	CountryCode   string  `json:"country_code"`
	Wind          float32 `json:"wind_spd"`
	WindDir       string  `json:"wind_cdir_full"`
	AppTemp       float32 `json:"app_temp"`
	WeatherStatus `json:"weather"`
}

//WeatherData struct
type WeatherData struct {
	Result []Weather `json:"data"`
}

//WeatherStatus struct
type WeatherStatus struct {
	Description string `json:"description"`
}

//Reminder has the information needed to survive a reboot of the bot:
type Reminder struct {
	DocID        string //Id from DATABASE
	UserID       string //UserID
	Message      string //Message to be reminded
	Date         string //Date at which is shall be reminded
	ChannelID    string //The channel the reminder was assigned too
	PulledFromDB bool   //If it was pulled from DB (bot was rebooted) it will be true
	toBeDeleted  bool   //Used for deleting if the reminder is sent (Has to be here)
}

//Bill struct
type Bill struct {
	ID    string
	Price string
	Type  string
}

//StartTime Used as start time from when the bot was booted:
var StartTime time.Time = time.Now()

//FirestoreDatabase struct
type FirestoreDatabase struct {
	CollectionName string
	Ctx            context.Context
	Client         *firestore.Client
}

//WebhookRegistration struct
type WebhookRegistration struct {
	ID        string `json:"ID"`
	URL       string `json:"url"`
	Event     string `json:"event"`
	Timestamp string `json:"Timestamp"`
}

// ConfigStruct struct:
type ConfigStruct struct {
	Token               string `json:"Token"`
	BotPrefix           string `json:"BotPrefix"`
	FirestoreCredential string `json:"FirestoreCredential"`
}

//Diagnostic holds diagnostics values
type Diagnostic struct {
	Uptime       int
	DBEmoji      string
	WeatherEmoji string
	WetherCode   int
	FoxEmoji     string
	FoxCode      int
}
