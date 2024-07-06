package models

import (
	"time"
)

type Comment struct {
	OwnerID   string      `json:"ownerID"` 
	Content   string    `json:"text"`
	Date      time.Time `json:"date"`
}
