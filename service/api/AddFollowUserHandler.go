package api

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) AddFollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract FollowerID from Authorization header
	FollowerID := r.Header.Get("Authorization")
	if FollowerID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Extract FollowingID from path parameter
	FollowingID := ps.ByName("userID")
	if FollowingID == "" {
		http.Error(w, "Missing userID parameter", http.StatusBadRequest)
		return
	}

	// Add follower relationship in the database
	err := rt.db.AddFollower(FollowerID, FollowingID)
	if err != nil {
		switch err.Error() {
		case "already followed":
			w.WriteHeader(http.StatusCreated)
			response := map[string]string{"message": "User is already being followed"}
			json.NewEncoder(w).Encode(response)
			if err != nil {
				http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
				log.Printf("Failed to encode response: %w", err)
			}
		case "user not found":
			http.Error(w, "Target user does not exist", http.StatusBadRequest)
		default:
			http.Error(w, "Failed to follow user", http.StatusInternalServerError)
		}
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
}
