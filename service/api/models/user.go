package models

type User struct {
    UserID   string    `db:"userID"`
    Username string `db:"username"`
    //Name     string `db:"name"`
	//Surname  string `db:"surname"`
	//Password string `db:"password"`
}

type Profile struct {
    UserID        string    `db:"userID"`
    Username      string  `db:"username"`
	FollowerList  []string  `json:"followerList"`
    FollowingList []string  `json:"followingList"`
    BanList       []string  `json:"banList"`
    PhotoList     []Picture `json:"photoList"`
}

