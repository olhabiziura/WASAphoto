package database

import (
	"errors"
)

func (db *appdbimpl) GetFeed(userID string) ([]string, error) {
	var feed []string

	rows, err := db.c.Query("SELECT pictureID FROM pictures WHERE ownerID=?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pictureID string
		err := rows.Scan(&pictureID)
		if err != nil {
			return nil, err
		}
		feed = append(feed, pictureID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Check if feed is empty and return an empty slice instead of an error
	if len(feed) == 0 {
		return []string{}, nil
	}

	return feed, nil
}
