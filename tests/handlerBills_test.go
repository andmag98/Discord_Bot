package tests

import (
	"project/src"
	"testing"
)

//TestRegisterBill tests functionality ragarding bills
func TestRegisterBill(t *testing.T) {
	src.ReadConfig("../config.json")
	err := src.InitDatabase("../")
	if err != nil {
		t.Error(err)
	}
	src.Start()
	src.DB.CollectionName = "TestID"
	id := "TestID"
	name := "TestName"

	text, err := src.RegisterBill(id, name, "100", "other")
	if err != nil {
		t.Error(err)
	}
	if text != ":white_check_mark: `I registered other for 100kr in TestName database!`" {
		t.Error(err)
	}
	textFail, err := src.RegisterBill(id, name, "abc", "50")
	if err != nil {
		t.Error(err)
	}
	if textFail != ":exclamation: `You wrote the price or type wrong!`" {
		t.Error(err)
	}

	message := src.SumAllHandler(id, name)
	if message == ":exclamation: `I was not able to get your bills from database!`" {
		t.Error(err)
	}
	if message != ":dollar: `TestName's total is: 100kr.`" {
		t.Error(err)
	}

	sumTypeMessage := src.SumTypeHandler(id, name, "other")
	if sumTypeMessage == ":exclamation: `I was not able to get your bills from database!`" {
		t.Error(err)
	}
	if sumTypeMessage != ":dollar: `TestName's total for other is: 100kr.`" {
		t.Error(err)
	}

	file, filemessage := src.DiagramHandler(id)
	if filemessage == ":exclamation: `I was not able to get your bills from database!`" {
		t.Error(err)
	}
	if file == nil {
		t.Error(err)
	}

	totalMessage := src.TotalHandler(id)
	if totalMessage == ":exclamation: `I was not able to get your bills from database!`" {
		t.Error(err)
	}

	everyBill, err := src.ReturnAllBills(id)
	if err != nil {
		t.Error("Cant find any bills")
	}

	if everyBill[0].ID != name && everyBill[0].Price != "100" && everyBill[0].Type != "other" {
		t.Errorf("Could not find in database")
	}

	//deleting everything from the database
	for j := range everyBill {
		err = src.DeleteBill(everyBill[j].ID)
		if err != nil {
			t.Errorf("ERROR: removal of Bills\n%v\n", err)
		}
	}
}
