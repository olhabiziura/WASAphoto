package database


func (db *appdbimpl) GetBan(userID string) ([]string, error) {
	var banList []string

	rows, err := db.c.Query("SELECT banID FROM ban WHERE userID=?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var banID string
		err := rows.Scan(&banID)
		if err != nil {
			return nil, err
		}
		banList = append(banList, banID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Check if banList is empty and return an empty slice instead of an error
	if len(banList) == 0 {
		return []string{}, nil
	}

	return banList, nil
}

func (db *appdbimpl) GetIfBan(userID, bannerID string) (bool, error) {
	// Adjusted query to properly use SELECT * FROM and fixed syntax with AND
	query := "SELECT 1 FROM ban WHERE banID=? AND userID=?"

	// Execute the query
	rows, err := db.c.Query(query, bannerID, userID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// Check if any rows are returned
	if rows.Next() {
		return true, nil
	}

	return false, nil
}
