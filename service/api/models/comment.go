package models

import (
	"time"
)

type Comment struct {
	CommentID string    `json:commentID`
	OwnerID   string    `json:"ownerID"`
	Username  string    `json:username`
	Content   string    `json:"text"`
	Date      time.Time `json:"date"`
	//OwnerUsername string `json:"ownerID"`
}
