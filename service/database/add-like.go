package database

// AddComment inserts a comment into the database.
func (db *appdbimpl) AddLike(userID, pictureID string) error {
	_, err := db.GetDB().Exec("INSERT INTO likes (userID, pictureID) VALUES (?, ?);", userID, pictureID)
	if err != nil {
		return err
	}
	return nil
}
