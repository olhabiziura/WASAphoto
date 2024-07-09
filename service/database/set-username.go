package database

import (
	"database/sql"
	"fmt"
)

// SetUsername updates the username of a user identified by userID to a new username.
func (db *appdbimpl) SetUsername(userID, newUsername string) error {
	// Check if the new username is already taken
	var existingUserID int64
	err := db.GetDB().QueryRow("SELECT userID FROM users WHERE username = ? LIMIT 1", newUsername).Scan(&existingUserID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking existing username: %v", err)
	}
	if existingUserID != 0 {
		return fmt.Errorf("username %s is already taken", newUsername)
	}

	// Update the username for the given userID
	_, err = db.GetDB().Exec("UPDATE users SET username = ? WHERE userID = ?", newUsername, userID)

	if err != nil {
		return fmt.Errorf("error updating username: %v", err)
	}

	return nil
}
