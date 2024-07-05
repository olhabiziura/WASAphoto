package database

import (
	
	"errors"
	"fmt"
)

var ErrAlreadyFollowed = errors.New("already followed")

// AddFollower inserts a follower relationship into the database.
func (db *appdbimpl) AddFollower(followerID, followingID string) error {
	fmt.Println(followerID, followingID)

	// Check if the user is already being followed
	var exists bool
	err := db.GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM followers WHERE followerID = ? AND followingID = ?)", followerID, followingID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return ErrAlreadyFollowed
	}

	// Insert the new follower relationship
	_, err = db.GetDB().Exec("INSERT INTO followers (followerID, followingID) VALUES (?, ?)", followerID, followingID)
	if err != nil {
		return err
	}

	return nil
}




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
