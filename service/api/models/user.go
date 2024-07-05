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
    Name          string  `db:"name"`
	Surname       string  `db:"surname"`
	FollowerList  []User  `json:"followerList"`
    FollowingList []User  `json:"followingList"`
    BanList       []User  `json:"banList"`
    PhotoList     []Picture `json:"photoList"`
}

