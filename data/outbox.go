package data

import (
	"fmt"
	"github.com/aasumitro/mego-worker/helper"
)

type Outbox struct {
	DestinationNumber string
	TextDecoded       string
	CreatorID         string
	Class             string
}

func StoreOutboxMessage(outbox Outbox) error {
	//load config
	config := helper.GetDataConfig()
	// call database connection function
	db, _ := DBConnection(config.MySQL)
	//defer the close till after the main function has finished executing
	defer db.Close()
	// data query insert
	query := fmt.Sprintf(
		"INSERT INTO Outbox (DestinationNumber, TextDecoded) VALUES ('%s', '%s')",
		outbox.DestinationNumber,
		outbox.TextDecoded)
	// perform a db.Query insert
	insert, err := db.Query(query)
	// if there is an error inserting, handle it
	helper.CheckError(err, "Failed insert new record")
	// be careful deferring Queries if you are using transactions
	// defer the close till after the main function has finished executing
	defer insert.Close()

	if err != nil {
		return err
	}

	return nil
}

func StoreSentMessage(outbox Outbox) error {
	//load config
	config := helper.GetDataConfig()
	// call database connection function
	db, _ := DBConnection(config.MySQL)
	//defer the close till after the main function has finished executing
	defer db.Close()
	// data query insert
	query := fmt.Sprintf(
		"INSERT INTO sentitems (DeliveryDateTime, Text, UDH, SenderID, CreatorID, DestinationNumber, TextDecoded) "+
			"VALUES ('%s', '%s')",
		outbox.DestinationNumber,
		outbox.TextDecoded)
	// perform a db.Query insert
	insert, err := db.Query(query)
	// if there is an error inserting, handle it
	helper.CheckError(err, "Failed insert new record")
	// be careful deferring Queries if you are using transactions
	// defer the close till after the main function has finished executing
	defer insert.Close()

	if err != nil {
		return err
	}

	return nil
}
