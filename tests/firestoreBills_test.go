package tests

import (
	"project/src"
	"strconv"
	"testing"
)

//TestFirestoreDatabase tests the database's bill functions
func TestFirestoreDatabase(t *testing.T) {

	src.ReadConfig("../config.json")
	err := src.InitDatabase("../")
	if err != nil {
		t.Errorf("Could init test database!")
	}
	defer src.DB.Client.Close()
	//testing on a test collection database
	collectionName := "testDB"
	src.DB.CollectionName = "testDB"

	bill := src.Bill{ID: "1", Price: "100", Type: "Clothes"}
	err = src.Save(&bill, collectionName)
	if err != nil {
		t.Errorf("Could not save to database!")
	}
	res, err := src.FindID(bill.ID)
	if err != nil {
		t.Errorf("Could not find in database")
	}

	if res[0].ID != bill.ID {
		t.Errorf("Bill.ID does not match!")
	}
	if res[0].Price != bill.Price {
		t.Errorf("Bill.Prise does not match!")
	}
	if res[0].Type != bill.Type {
		t.Errorf("Bill.Type does not match!")
	}

	// we should clean-up
	err = src.DeleteBill(res[0].ID)
	if err != nil {
		t.Errorf("ERROR: removal of Bill\n%v\n", err)
	}
	resAfterRemove, err := src.FindID(res[0].ID)
	if err != nil {
		t.Error("FindID function has failed!")
	}
	// THERE MUST BE AN ERROR
	if len(resAfterRemove) != 0 {
		t.Error("FindID has not failed for deleted document!")
	}
	//puts testdata in the database
	var bills [6]src.Bill
	for i := 0; i <= 5; i++ {
		bill := src.Bill{ID: strconv.Itoa(i), Price: "100" + strconv.Itoa(i),
			Type: "Test" + strconv.Itoa(i)}
		err := src.Save(&bill, collectionName)
		if err != nil {
			t.Error("Save has failed to add document!")
		}
		bills[i] = bill
	}
	//gets all the data from the database
	everyBill, err := src.ReturnAllBills(collectionName)
	if err != nil {
		t.Error("Cant find any bills")
	}
	//if the length of testdata is not the same as the database returned
	if len(everyBill) != len(bills) {
		t.Error("Did not return every registerd bill")
	}

	//deleting everything from the database
	for j := range everyBill {
		err = src.DeleteBill(everyBill[j].ID)
		if err != nil {
			t.Errorf("ERROR: removal of Bills\n%v\n", err)
		}
	}
}
