package database

import (
	"database/sql"
	"fmt"
	"errors"
)

// DeletePhoto deletes a photo record from the database and returns the file path of the deleted photo.
func (db *appdbimpl) DeletePhoto(userID, photoID string) (string, error) {
	// Query to fetch photo address and delete the photo record
	var filePath string
	err := db.GetDB().QueryRow("SELECT address FROM pictures WHERE ownerID = ? AND pictureID = ?", userID, photoID).Scan(&filePath)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("photo not found")
		}
		return "", fmt.Errorf("error retrieving photo: %w", err)
	}
	// Execute delete query
	_, err = db.GetDB().Exec("DELETE FROM pictures WHERE ownerID = ? AND pictureID = ?", userID, photoID)
	if err != nil {
		return "", fmt.Errorf("error deleting photo: %w", err)
	}

	return filePath, nil
}
