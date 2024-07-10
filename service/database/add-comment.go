package database

import (
	"fmt"
)

// AddComment inserts a comment into the database and returns the commentID of the newly created comment.
func (db *appdbimpl) AddComment(userID, pictureID, content string) (string, error) {
	result, err := db.GetDB().Exec("INSERT INTO comments (ownerID, pictureID, text) VALUES (?, ?, ?);", userID, pictureID, content)
	if err != nil {
		return "", err
	}

	commentID, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", commentID), nil
}
