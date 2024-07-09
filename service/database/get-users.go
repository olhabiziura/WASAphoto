package database

import (
	"fmt"
	"wasaphoto/service/api/models"
)

// GetUsersByUsernamePrefix queries the database for users whose usernames start with the given prefix
func (db *appdbimpl) GetUsers(usernamePrefix string) ([]models.User, error) {
	rows, err := db.c.Query("SELECT userID, username FROM users WHERE username LIKE ? ORDER BY username ASC", usernamePrefix+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.Username); err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over user rows: %w", err)
	}

	return users, nil
}
