package database

// AddComment inserts a comment into the database.
func (db *appdbimpl) AddComment(userID, pictureID , content string) error {
	_, err := db.GetDB().Exec("INSERT INTO comments (ownerID, pictureID, text) VALUES (?, ?, ?);", userID, pictureID, content)
	if err != nil {
		return err
	}
	return nil
}
