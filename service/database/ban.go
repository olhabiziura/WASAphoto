package database

import (
	"errors"
	"fmt"
)

var ErrAlreadyBanned = errors.New("already banned")

// AddBan adds a user to the banned list in the database.
func (db *appdbimpl) AddBan(UserID, BanID string) error {
	fmt.Println(UserID, BanID)

	// Check if the user is already banned
	var exists bool
	err := db.GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM ban WHERE userID = ? AND banID = ?)", UserID, BanID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return ErrAlreadyBanned
	}

	// Insert the new ban
	_, err = db.GetDB().Exec("INSERT INTO ban (userID, banID) VALUES (?, ?)", UserID, BanID)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) DeleteBan(UserID, BannedID string) error {
	fmt.Println(UserID, BannedID)

	// Execute the SQL delete statement
	_, err := db.GetDB().Exec("DELETE FROM ban WHERE userID = ? AND banID = ?", UserID, BannedID)
	if err != nil {
		return err
	}

	return nil
}
