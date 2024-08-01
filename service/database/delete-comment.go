package database

func (db *appdbimpl) DeleteComment(userID, pictureID, commentID string) error {
	// Execute the SQL delete statement
	_, err := db.c.Exec("DELETE FROM comments WHERE pictureID = ? AND commentID = ?", pictureID, commentID)
	if err != nil {
		return err
	}

	return nil
}
