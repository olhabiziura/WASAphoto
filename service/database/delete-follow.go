package database

import (

	"errors"
	"fmt"
)

var ErrNotFollowing = errors.New("not following")

// DeleteFollower removes a follower relationship from the database.
func (db *appdbimpl) DeleteFollower(followerID, followingID string) error {
	fmt.Println(followerID, followingID)

	// Check if the user is currently following
	var exists bool
	err := db.GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM followers WHERE followerID = ? AND followingID = ?)", followerID, followingID).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return ErrNotFollowing
	}

	// Execute the SQL delete statement
	_, err = db.GetDB().Exec("DELETE FROM followers WHERE followerID = ? AND followingID = ?", followerID, followingID)
	if err != nil {
		return err
	}

	return nil
}
