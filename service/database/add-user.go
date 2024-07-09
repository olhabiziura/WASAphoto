package database

import (
	"fmt"
)

// AddUser inserts a new user into the database and returns the user_id.
func (db *appdbimpl) AddUser(username string) (int64, error) {
	// Execute the insert query with parameterized values to prevent SQL injection
	result, err := db.c.Exec("INSERT INTO users (username) VALUES (?)", username)
	fmt.Println(result, err)
	if err != nil {
		return 0, err
	}

	// Get the ID of the last inserted user
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
