package database

import (
	"fmt"
	"wasaphoto/service/api/models"
)

func (db *appdbimpl) GetComments(pictureID string) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := db.c.Query("SELECT commentID, ownerID, text FROM comments WHERE pictureID=?", pictureID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var ownerID string
		var content string
		var commentID string
		err := rows.Scan(&commentID, &ownerID, &content)
		if err != nil {
			return nil, err
		}
		username, err := db.GetUsername(ownerID)
		// Populate Comment struct
		comment.OwnerID = ownerID
		comment.Content = content
		comment.CommentID = commentID
		comment.Username = username
		// Append to comments slice
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return an empty slice if no comments are found
	return comments, nil
}
