package database

import (
	"fmt"
)

func (db *appdbimpl) DeleteBan(UserID, BannedID string) error {
	fmt.Println(UserID, BannedID)

	// Execute the SQL delete statement
	_, err := db.GetDB().Exec("DELETE FROM ban WHERE userID = ? AND banID = ?", UserID, BannedID)
	if err != nil {
		return err
	}

	return nil
}
