package src

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

//DB database
var DB = FirestoreDatabase{}

//InitDatabase function initializes the database
func InitDatabase(filepath string) error {
	DB.Ctx = context.Background()
	sa := option.WithCredentialsFile(filepath + FirestoreCredential)
	app, err := firebase.NewApp(DB.Ctx, nil, sa)
	if err != nil {
		return errors.Wrap(err, "Something went wrong with Initdatabase()")
	}
	DB.Client, err = app.Firestore(DB.Ctx)
	if err != nil {
		return errors.Wrap(err, "Something went wrong with Initdatabase()")
	}
	return nil
}

//Save function saves a bill in the database
func Save(b *Bill, collectionName string) error {
	ref := DB.Client.Collection(collectionName).NewDoc()
	b.ID = ref.ID
	_, err := ref.Set(DB.Ctx, b)
	if err != nil {
		fmt.Println("ERROR saving bill to Firestore DB: ", err)
		return errors.Wrap(err, "Error in FirebaseDatabase.Save()")
	}
	return nil
}

//GetData from database function
func GetData(w http.ResponseWriter) error {
	iter := DB.Client.Collection(DB.CollectionName).Documents(DB.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		err2 := json.NewEncoder(w).Encode(doc.Data())
		if err2 != nil {
			return err2
		}
	}
	return nil
}

//ReturnAllBills returns all the bills from a given collection in the database
func ReturnAllBills(collectionName string) ([]Bill, error) {
	var bills []Bill

	iter := DB.Client.Collection(collectionName).Documents(DB.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err == iterator.Done {
			return nil, err
		}
		bill := Bill{}
		e := doc.DataTo(&bill)
		if e != nil {
			return nil, e
		}
		bills = append(bills, bill)
	}
	return bills, nil
}

//FindID function find a bill with id
func FindID(id string) ([]Bill, error) {
	iter := DB.Client.Collection(DB.CollectionName).Where("ID", "==", id).Documents(DB.Ctx)
	var bills = []Bill{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Error when querying for bill: %v\n%v\n", id, err)
		}

		bill := Bill{}
		err = doc.DataTo(&bill)
		if err != nil {
			fmt.Println("Error when converting retrieved document to Bill struct: ", err)
		}
		bills = append(bills, bill)
	}
	return bills, nil
}

//DeleteBill deletes a bill from the database
func DeleteBill(id string) error {
	_, err := DB.Client.Collection(DB.CollectionName).Doc(id).Delete(DB.Ctx)
	if err != nil {
		fmt.Printf("ERROR deleting bill (%v) from Firestore DB: %v\n", id, err)
		return errors.Wrap(err, "Error in FirebaseDatabase.Delete()")
	}
	return nil
}
