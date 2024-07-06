package database

import (
	"wasaphoto/service/api/models"
)

func (db *appdbimpl) GetFeed(userID string) ([]models.Picture, error) {
	var feed []models.Picture

	rows, err := db.c.Query("SELECT pictureID, ownerID, date, address FROM pictures WHERE ownerID=?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pic models.Picture
		err := rows.Scan(&pic.PictureID, &pic.OwnerID, &pic.Date, &pic.Address)
		if err != nil {
			return nil, err
		}

		// Call GetLikes to fetch the list of users who liked this picture
		pic.Likes, err = db.GetLikes(pic.PictureID)
		if err != nil {
			return nil, err
		}

		// Call GetComments to fetch comments for this picture
		pic.Comments, err = db.GetComments(pic.PictureID)
		if err != nil {
			return nil, err
		}

		feed = append(feed, pic)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Check if feed is empty and return an empty slice instead of an error
	if len(feed) == 0 {
		return []models.Picture{}, nil
	}

	return feed, nil
}
