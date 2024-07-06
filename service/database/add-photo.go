package database

import (
	
	"fmt"
	"time"
)

// AddPhoto inserts a record into the pictures table with details of the added photo
// and returns the ID of the newly inserted picture.
func (db *appdbimpl) AddPhoto(ownerID, address string, date time.Time, description string) (int64, error) {
	var pictureID int64

	// Depending on the SQL dialect, this part may vary
	query := "INSERT INTO pictures (ownerID, address, date, description) VALUES (?, ?, ?, ?)"

	result, err := db.GetDB().Exec(query, ownerID, address, date, description)
	if err != nil {
		return 0, fmt.Errorf("error inserting photo record: %v", err)
	}

	// For SQL databases that support LastInsertId (e.g., SQLite, MySQL)
	pictureID, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error retrieving last insert ID: %v", err)
	}

	return pictureID, nil
}
