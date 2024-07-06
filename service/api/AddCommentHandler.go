package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AddCommentHandler handles adding comments using the net/http package
func (rt *_router) AddComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get UserID from Authorization header
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Get PictureID from path parameter
	pictureID := ps.ByName("PictureId")
	if pictureID == "" {
		http.Error(w, "Missing pictureID parameter", http.StatusBadRequest)
		return
	}

	// Read comment content from request body
	var requestBody map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	content, ok := requestBody["text"]
	if !ok || content == "" {
		http.Error(w, "Comment text cannot be empty", http.StatusBadRequest)
		return
	}

	// Insert comment into the database using the AddComment function
	if err := rt.db.AddComment(userID, pictureID, content); err != nil {
		http.Error(w, "Failed to add a comment", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Comment added successfully"})
}
