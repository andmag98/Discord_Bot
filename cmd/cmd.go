package cmd

import (
	"fmt"
	"project/src"
)

//Run runs the programs, acts as a pseudo main
func Run() {
	src.ReadConfig("config.json")
	err := src.InitDatabase("./")
	if err != nil {
		fmt.Println(err)
	}
	err = src.ReminderInit("./")
	if err != nil {
		fmt.Println(err)
	}

	src.Start()

	<-make(chan struct{})
	defer src.RemindClient.Close()
	defer src.DB.Client.Close()
}
