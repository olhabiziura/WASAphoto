package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"

)

func (rt *_router)DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Extract UserID from request headers
	UserID := r.Header.Get("Authorization")
	if UserID == "" {
		http.Error(w, "Missing user_id in headers", http.StatusBadRequest)
		return
	}

	// Extract PictureID from URL query parameters
	PictureID := r.URL.Query().Get("PictureID")
	if PictureID == "" {
		http.Error(w, "Missing PictureID in query parameters", http.StatusBadRequest)
		return
	}
    
	CommentID := r.URL.Query().Get("CommentID")
	if CommentID == "" {
		http.Error(w, "Missing CommentID in query parameters", http.StatusBadRequest)
		return
	}


	// Delete the comment from the database
	err := rt.db.DeleteComment(UserID, PictureID, CommentID)
	if err != nil {
		http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Comment deleted successfully"}
	json.NewEncoder(w).Encode(response)
}
