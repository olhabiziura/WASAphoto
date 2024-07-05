package models

import (
	"time"
)

type Comment struct {
	CommentID string    `json:"commentID"`
	PictureID string       `db:"pictureID"` 
	OwnerID   User      `json:"ownerID"` 
	Content   string    `json:"text"`
	Date      time.Time `json:"date"`
}
