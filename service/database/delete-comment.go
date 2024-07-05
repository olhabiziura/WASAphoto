package database

import (
	"database/sql"
)


func getCommentOwner(db *sql.DB, userID, commentID string) (bool, error) {
    var ownerID string
    err := db.QueryRow("SELECT ownerID FROM comments WHERE commentID = ?", commentID).Scan(&ownerID)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil // Comment with commentID not found
        }
        return false, err // Other errors
    }
    
    return ownerID == userID, nil
}


func (db *appdbimpl) DeleteComment(userID, pictureID, commentID string) error {
    // Check if the user owns the comment
    isOwner, err := getCommentOwner(db.c, userID, commentID)
    if err != nil {
        return err
    }
    
    if !isOwner {
        return sql.ErrNoRows // Return appropriate error when user does not own the comment
    }
    
    // Execute the SQL delete statement
    _, err = db.c.Exec("DELETE FROM comments WHERE userID = ? AND pictureID = ? AND commentID = ?", userID, pictureID, commentID)
    if err != nil {
        return err
    }
    
    return nil
}