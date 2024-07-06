package models

import (
	"time"
)

type Picture struct {
	PictureID string   `db:"pictureID"`
	OwnerID   string    `db:"ownerID"`
	Username string `json:username`
	Likes     []string `json:"likes"`
	Comments []Comment `json:comments`
	Date    time.Time `db:"date"`
	Address string `db:"address"`
	Image     string    `json:"image"`
}
