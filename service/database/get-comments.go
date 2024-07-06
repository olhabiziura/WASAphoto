package database

import (

	"wasaphoto/service/api/models"
	"fmt"
)

func (db *appdbimpl) GetComments(pictureID string) ([]models.Comment, error) {
	var comments []models.Comment
	fmt.Println("jonljnl")
	rows, err := db.c.Query("SELECT ownerID, text FROM comments WHERE pictureID=?", pictureID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var ownerID string
		var content string


		err := rows.Scan(&ownerID, &content)
		if err != nil {
			return nil, err
		}

		// Populate Comment struct
		comment.OwnerID = ownerID
		comment.Content = content


		// Append to comments slice
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return an empty slice if no comments are found
	return comments, nil
}
