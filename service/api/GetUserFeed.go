package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

)

func (rt *_router) GetUserFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get UserID from Authorization header
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Get feed for the user
	feed, err := rt.db.GetFeed(userID)
	if err != nil {
		switch err.Error() {
		case "invalid user_id":
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
		default:
			http.Error(w, "Failed to get feed", http.StatusInternalServerError)
		}
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Marshal feed to JSON and send response
	json.NewEncoder(w).Encode(feed)
}

