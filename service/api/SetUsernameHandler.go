package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SetUsernameHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get UserID from Authorization header
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Decode JSON request body to get the new username
	var newUsername struct {
		Username string `json:"username"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUsername); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Call the database function to set the new username
	err := rt.db.SetUsername(userID, newUsername.Username)
	if err != nil {
		switch err.Error() {
		case "username %s is already taken":
			http.Error(w, fmt.Sprintf("Username %s is already taken", newUsername.Username), http.StatusBadRequest)
		default:
			http.Error(w, fmt.Sprintf("Failed to change username: %v", err), http.StatusInternalServerError)
		}
		return
	}

	// Send success response
	w.WriteHeader(http.StatusNoContent)
}
