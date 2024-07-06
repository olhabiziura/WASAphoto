package database

import (
	"errors"
)

func (db *appdbimpl) GetFollowing(userID string) ([]string, error) {
	var followingList []string

	rows, err := db.c.Query("SELECT FollowingID FROM followers WHERE FollowerID=?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var followingID string
		err := rows.Scan(&followingID)
		if err != nil {
			return nil, err
		}
		followingList = append(followingList, followingID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Check if followingList is empty and return an empty slice instead of an error
	if len(followingList) == 0 {
		return []string{}, nil
	}

	return followingList, nil
}
