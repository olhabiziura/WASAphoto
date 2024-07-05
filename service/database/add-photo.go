package database

import (

	"fmt"
	"time"
)

// AddPhoto inserts a record into the pictures table with details of the added photo.
func (db *appdbimpl) AddPhoto(ownerID, address string, date time.Time, description string) error {
	_, err := db.GetDB().Exec("INSERT INTO pictures (ownerID, address, date, description) VALUES (?, ?, ?, ?)",
		ownerID, address, date, description)
	if err != nil {
		return fmt.Errorf("error inserting photo record: %v", err)
	}
	return nil
}
