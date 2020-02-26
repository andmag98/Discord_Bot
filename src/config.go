package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// Token variable
	Token string

	// BotPrefix variable
	BotPrefix string

	//FirestoreCredential file:
	FirestoreCredential string

	config *ConfigStruct
)

// ReadConfig reads all config variables from the config.json file
func ReadConfig(filePath string) {
	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	FirestoreCredential = config.FirestoreCredential
}
