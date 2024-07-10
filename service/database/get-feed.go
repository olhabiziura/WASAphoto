package database

import (
	"log"
	"strconv"
	"time"
	"wasaphoto/service/api/models"
)

func (db *appdbimpl) GetFeed(userID string) ([]models.Picture, error) {
	var feed []models.Picture

	rows, err := db.c.Query("SELECT pictureID, ownerID, date, address FROM pictures WHERE ownerID=?", userID)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pic models.Picture

		var pictureID string
		var ownerID int
		var date string // Use string instead of time.Time
		var address string

		err := rows.Scan(&pictureID, &ownerID, &date, &address)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}

		// Parse the date string into a time.Time variable
		parsedDate, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", date) // Adjust the layout according to your date format
		if err != nil {
			log.Printf("Failed to parse date: %v", err)
			return nil, err
		}

		// Populate the pic struct
		pic.PictureID = pictureID
		pic.OwnerID = strconv.Itoa(ownerID)
		pic.Date = parsedDate
		pic.Address = address

		username, err := db.GetUsername(strconv.Itoa(ownerID))
		if err != nil {
			log.Printf("Failed to get username: %v", err)
			return nil, err
		}
		pic.Username = username

		// Call GetLikes to fetch the list of users who liked this picture
		pic.Likes, err = db.GetLikes(pic.PictureID)
		if err != nil {
			log.Printf("Failed to get likes: %v", err)
			return nil, err
		}

		// Call GetComments to fetch comments for this picture
		pic.Comments, err = db.GetComments(pic.PictureID)
		if err != nil {
			log.Printf("Failed to get comments: %v", err)
			return nil, err
		}

		feed = append(feed, pic)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		return nil, err
	}

	// Check if feed is empty and return an empty slice instead of an error
	if len(feed) == 0 {
		return []models.Picture{}, nil
	}

	return feed, nil
}
