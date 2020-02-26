package tests

import (
	"errors"
	"project/src"
	"testing"
	"time"
)

//TestNewRunningReminder
func TestNewRunningReminder(t *testing.T) {

	err := errors.New("Invalid function parameters. hours, minutes or seconds are below 0")
	err2 := errors.New("Invalid function parameters. channelID can't be empty")
	var tests = []struct {
		inputID            string
		inputMessage       string
		inputUser          string
		inputDiscriminator string
		inputsek           int
		inputmin           int
		inputhours         int

		expectedError error
	}{
		{"123", "", "", "", 0, 0, 0, nil},
		{"123", "", "", "", -1, 0, 0, err},
		{"123", "", "", "", 0, -1, 0, err},
		{"123", "", "", "", 0, 0, -1, err},
		{"", "", "", "", 0, 0, 0, err2},
	}

	for i, test := range tests {
		_, e := src.NewRunningReminder(test.inputID, test.inputMessage, test.inputUser,
			test.inputDiscriminator, test.inputsek, test.inputmin, test.inputhours)
		if e != nil && test.expectedError != nil {
			if e.Error() != test.expectedError.Error() {
				t.Errorf("Test "+string(i)+" failed: %v %v %v %v %v %v %v inputed, %v expected, received: %v", test.inputID, test.inputMessage, test.inputUser,
					test.inputDiscriminator, test.inputsek, test.inputmin, test.inputhours, test.expectedError, e)
			}
		} else {
			if e != test.expectedError {
				t.Errorf("Test "+string(i)+" failed: %v %v %v %v %v %v %v inputed, %v expected, received: %v", test.inputID, test.inputMessage, test.inputUser,
					test.inputDiscriminator, test.inputsek, test.inputmin, test.inputhours, test.expectedError, e)
			}
		}

	}
}

//TestNewRunningReminderDB
func TestNewRunningReminderDB(t *testing.T) {
	var startTime time.Time = time.Now()
	err := errors.New("Function called without all using both parameters")
	//err2 := errors.New("Function called with an invalid, empty string in perameters")
	var tests = []struct {
		inputMessage  string
		inputAtUser   string
		inputTime     *time.Time
		expectedError error
	}{
		{"123", "", &startTime, nil},
		{"", "", &startTime, err},
		{"", "user", &startTime, err},
		{"123", "user", &startTime, nil},
	}

	for i, test := range tests {
		_, e := src.NewRunningReminderDB(test.inputMessage, test.inputAtUser, test.inputTime)
		if e != nil && test.expectedError != nil {
			if e.Error() != test.expectedError.Error() {
				t.Errorf("Test failed: %v %v %v inputed, %v expected, received: %v", test.inputMessage, test.inputAtUser, test.inputTime, test.expectedError, e)
			}
		} else {
			if e != test.expectedError {
				t.Errorf("Test %v failed: %v %v %v inputed, %v expected, received: %v", i, test.inputMessage, test.inputAtUser, test.inputTime, test.expectedError, e)
			}
		}

	}
}
