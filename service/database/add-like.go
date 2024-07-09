package database

import (
	"database/sql"
	"errors"
)

// AddLike inserts a like into the database if it doesn't already exist.
func (db *appdbimpl) AddLike(userID, pictureID string) error {
	// Check if the like already exists
	exists, err := db.likeExists(userID, pictureID)
	if err != nil {
		return err
	}
	if exists {
		// Like already exists, return nil indicating success (no error)
		return nil
	}

	// If not exists, insert the like
	_, err = db.GetDB().Exec("INSERT INTO likes (userID, pictureID) VALUES (?, ?);", userID, pictureID)
	if err != nil {
		return err
	}
	return nil
}

// Helper function to check if a like exists
func (db *appdbimpl) likeExists(userID, pictureID string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM likes WHERE userID = ? AND pictureID = ? LIMIT 1);"
	err := db.GetDB().QueryRow(query, userID, pictureID).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// If no rows found, return false (like does not exist)
			return false, nil
		}
		// Return any other errors encountered during query
		return false, err
	}
	// Return true if like exists
	return exists, nil
}
