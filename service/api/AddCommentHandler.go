package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// AddCommentHandler handles adding comments using net/http package
func (rt *_router) AddComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	UserID := r.Header.Get("Authorization")
	if UserID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

    PictureID := r.URL.Query().Get("PictureID")

    // Decode comment content from header
    Content := r.Header.Get("Text")

    // Insert comment into the database using the AddComment function
    err := rt.db.AddComment(UserID, PictureID, Content)
    if err != nil {
        http.Error(w, "Failed to add a comment", http.StatusInternalServerError)
        return
    }

    // Respond with success message
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "comment added successfully"})
}
