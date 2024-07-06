package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetUsername(userID string) (string, error) {
	var username string
	err := db.c.QueryRow("SELECT username FROM users WHERE userID=?", userID).Scan(&username)
	return username, err
}
