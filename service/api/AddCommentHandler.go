package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
)

// AddCommentHandler handles adding comments using net/http package
func (rt *_router) AddComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get UserID from Authorization header
	UserID := r.Header.Get("Authorization")
	if UserID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Get PictureID from path parameter
	PictureID := ps.ByName("pictureID")
	if PictureID == "" {
		http.Error(w, "Missing pictureID parameter", http.StatusBadRequest)
		return
	}

	// Read comment content from request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Unmarshal JSON body to get the comment text
	var requestBody map[string]string
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	Content, ok := requestBody["text"]
	if !ok {
		http.Error(w, "Missing 'text' field in request body", http.StatusBadRequest)
		return
	}

	// Insert comment into the database using the AddComment function
	err = rt.db.AddComment(UserID, PictureID, Content)
	if err != nil {
		http.Error(w, "Failed to add a comment", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "comment added successfully"})
}
