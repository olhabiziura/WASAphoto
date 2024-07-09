package database

import (
	"fmt"
	"wasaphoto/service/api/models"
)

func (db *appdbimpl) GetFollowing(userID string) ([]models.User, error) {
	var followingList []models.User

	// Construct SQL query with a join to fetch followers and their usernames
	query := `
		SELECT u.userID, u.username
		FROM followers f
		JOIN users u ON f.FollowingID = u.userID
		WHERE f.FollowerID=?
	`

	rows, err := db.c.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying followers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UserID, &user.Username)
		if err != nil {
			return nil, fmt.Errorf("error scanning follower: %w", err)
		}
		followingList = append(followingList, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return followingList, nil
}
