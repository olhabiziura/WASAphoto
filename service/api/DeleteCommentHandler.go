package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"

)

func (rt *_router) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract UserID from request headers
	UserID := r.Header.Get("Authorization")
	if UserID == "" {
		http.Error(w, "Missing user_id in headers", http.StatusBadRequest)
		return
	}

	// Extract PictureID and CommentID from URL parameters
	PictureID := ps.ByName("PictureID")
	if PictureID == "" {
		http.Error(w, "Missing PictureID parameter", http.StatusBadRequest)
		return
	}

	CommentID := ps.ByName("CommentID")
	if CommentID == "" {
		http.Error(w, "Missing CommentID parameter", http.StatusBadRequest)
		return
	}

	// Delete the comment from the database
	err := rt.db.DeleteComment(UserID, PictureID, CommentID)
	if err != nil {
		// Check specific error conditions and return appropriate status codes
		switch err.Error() {
		case "comment not found":
			http.Error(w, "Comment not found", http.StatusBadRequest)
		case "comment does not belong to user":
			http.Error(w, "Comment does not belong to the user", http.StatusBadRequest)
		default:
			http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
		}
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Comment deleted successfully"}
	json.NewEncoder(w).Encode(response)
}
