package tests

import (
	"errors"
	"fmt"
	"project/src"
	"testing"
	"time"
)

func TestAddReminderToDB(t *testing.T) {
	var startTime time.Time = time.Now()
	myTestReminders := make([]src.Reminder, 0)

	err := errors.New("Invalid function parameters. channelID, m and remindDate has to be filled")
	var tests = []struct {
		inputChannelID string
		inputMessage   string
		inputAtUser    string
		inputTime      *time.Time
		expectedError  error
	}{
		{"", "", "", &startTime, err},
		{"", "", "123", &startTime, err},
		{"", "123", "", &startTime, err},
		{"", "123", "123", &startTime, err},
		{"123", "", "", &startTime, err},
		{"123", "", "123", &startTime, err},
		{"123", "123", "", &startTime, nil},
		{"123", "123", "123", &startTime, nil},
		{"123", "123", "123", nil, err},
		{"123", "123", "123", &startTime, nil},
		{"", "", "", &startTime, err},
	}
	for i, test := range tests {
		r, e := src.NewReminder(test.inputChannelID, test.inputMessage, test.inputAtUser, test.inputTime)

		if e == nil {
			if e != test.expectedError {
				t.Errorf("Test %v failed: %v %v %v inputed, %v expected, received: %v", i, test.inputMessage,
					test.inputAtUser, test.inputTime, test.expectedError, e)
			} else {
				myTestReminders = append(myTestReminders, *r)
			}
		} else {
			if e.Error() != test.expectedError.Error() {
				t.Errorf("Test %v failed: %v %v %v inputed, %v expected, received: %v", i, test.inputMessage, test.inputAtUser, test.inputTime, test.expectedError, e)
			}
		}

	}
	src.ReadConfig("../config.json")
	err = src.ReminderInit("../")
	if err != nil {
		t.Errorf("Test failed: no error expected when connecting to DB")
		fmt.Println(err)
	} else {
		defer src.RemindClient.Close()

		for i, testRemind := range myTestReminders {
			err = src.AddReminderToDB(&testRemind)
			if err != nil {
				t.Errorf("Test %v failed: no error expected when adding reminder to DB", i)
			}
		}
		myTestReminders, err = src.GetReminderFromDB()
		if err != nil {
			t.Errorf("Test failed: no error expected when getting reminder to DB")
		}

		for i, testRemind := range myTestReminders {
			err = testRemind.DeleteReminderFromDB()
			if err != nil {
				t.Errorf("Test %v failed: no error expected when Deleting reminder to DB", i)
			}
		}
	}
}
