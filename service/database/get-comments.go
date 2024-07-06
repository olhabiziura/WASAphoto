package database

import (
	"time"
	"wasaphoto/service/api/models"
)

func (db *appdbimpl) GetComments(pictureID string) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := db.c.Query("SELECT ownerID, text, date FROM comments WHERE pictureID=?", pictureID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var ownerID string
		var content string
		var date time.Time

		err := rows.Scan(&ownerID, &content, &date)
		if err != nil {
			return nil, err
		}

		// Populate Comment struct
		comment.OwnerID = ownerID
		comment.Content = content
		comment.Date = date

		// Append to comments slice
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}