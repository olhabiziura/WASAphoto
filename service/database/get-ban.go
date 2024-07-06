package database

import (
	"errors"
)

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
