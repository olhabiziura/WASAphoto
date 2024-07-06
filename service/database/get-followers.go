package database


func (db *appdbimpl) GetFollowers(userID string) ([]string, error) {
	var followerList []string

	rows, err := db.c.Query("SELECT FollowerID FROM followers WHERE FollowingID=?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var followerID string
		err := rows.Scan(&followerID)
		if err != nil {
			return nil, err
		}
		followerList = append(followerList, followerID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Check if followerList is empty and return an empty slice instead of an error
	if len(followerList) == 0 {
		return []string{}, nil
	}

	return followerList, nil
}
