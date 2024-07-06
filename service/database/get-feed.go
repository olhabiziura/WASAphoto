package database

import (
	"wasaphoto/service/api/models"
	"fmt"
	"time"
	"log"
	"strconv"
)

func (db *appdbimpl) GetFeed(userID string) ([]models.Picture, error) {
	var feed []models.Picture
	fmt.Println("enter GetFeed")
	rows, err := db.c.Query("SELECT pictureID, ownerID, date, address FROM pictures WHERE ownerID=?", userID)
	if err != nil {
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
			log.Fatal(err)
		}

		// Parse the date string into a time.Time variable
		parsedDate, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", date) // Adjust the layout according to your date format
		if err != nil {
			log.Fatal(err)
		}

		

		// Populate the pic struct
		pic.PictureID = pictureID
		pic.OwnerID = strconv.Itoa(ownerID)
		pic.Date = parsedDate
		pic.Address = address
		fmt.Println("lhblikb")
		
		username, err := db.GetUsername(strconv.Itoa(ownerID))
		if err != nil {
			return nil, err
		}
		pic.Username = username


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
	fmt.Println("exit get feedS")
	return feed, nil
}
