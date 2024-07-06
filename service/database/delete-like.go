package database

func (db *appdbimpl) DeleteLike(UserID, PictureID string) error {
    // Execute the SQL delete statement
    _, err := db.GetDB().Exec("DELETE FROM likes WHERE userID = ? AND pictureID = ?", UserID,PictureID)
    if err != nil {
        return err
    }
    return nil
}