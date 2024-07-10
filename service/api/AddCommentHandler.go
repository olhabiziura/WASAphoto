package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"wasaphoto/service/api/models"
)

// AddCommentHandler handles adding comments using the net/http package
func (rt *_router) AddComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var comment models.Comment
	// Get UserID from Authorization header
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Get PictureID from path parameter
	pictureID := ps.ByName("pictureID")
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
	commentID, err := rt.db.AddComment(userID, pictureID, content)
	if err != nil {
		http.Error(w, "Failed to add a comment", http.StatusInternalServerError)
		return
	}

	// Populate comment fields
	comment.OwnerID = userID
	comment.Content = content
	comment.CommentID = commentID

	// Respond with success message and the added comment
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Comment added successfully",
		"comment": comment,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
	}
}
