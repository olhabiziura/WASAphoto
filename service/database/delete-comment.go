package database



func (db *appdbimpl) DeleteComment(userID, pictureID, commentID string) error {
    // Execute the SQL delete statement
    _, err := db.c.Exec("DELETE FROM comments WHERE ownerID = ? AND pictureID = ? AND commentID = ?", userID, pictureID, commentID)
    if err != nil {
        return err
    }
    
    return nil
}
