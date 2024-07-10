package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) Login(username string) (int64, error) {
	// Query to check if the user exists
	query := "SELECT userID FROM users WHERE username = ?"
	var userID int64
	err := db.GetDB().QueryRow(query, username).Scan(&userID)

	// Handling errors and logic based on query results
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User does not exist, so add the user
			userID, err = db.AddUser(username)
			if err != nil {
				return 0, fmt.Errorf("failed to add user: %w", err)
			}
			return userID, nil
		}
		// Other error occurred during query execution
		return 0, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return userID, nil
}
