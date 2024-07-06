package database

// GetLikesByPictureID retrieves a list of userIDs who liked a specific pictureID
func (db *appdbimpl) GetLikes(pictureID string) ([]string, error) {
	var userIDs []string

	rows, err := db.c.Query("SELECT userID FROM likes WHERE pictureID=?", pictureID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userID string
		err := rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userIDs, nil
}
