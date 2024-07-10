package models

type User struct {
	UserID   string `db:"userID"`
	Username string `db:"username"`
}

type Profile struct {
	UserID        string    `db:"userID"`
	Username      string    `db:"username"`
	FollowerList  []User    `json:"followerList"`
	FollowingList []User    `json:"followingList"`
	BanList       []string  `json:"banList"`
	PhotoList     []Picture `json:"photoList"`
}
