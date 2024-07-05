package models

import (
	"time"
)

type Picture struct {
	PictureID string    `db:"pictureID"`
	OwnerID   string    `db:"ownerID"`
	Likes     []User `json:"likes"`
	Date    time.Time `db:"date"`
	Format   string  `db:"format"`
	Address string `db:"address"`
}
