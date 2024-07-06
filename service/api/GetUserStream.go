package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto/service/api/models"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetUserStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get UserID from Authorization header
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Get following list for the user
	followingList, err := rt.db.GetFollowing(userID)
	if err != nil {
		http.Error(w, "Failed to get following list", http.StatusInternalServerError)
		return
	}

	// Retrieve feed for each following (last 5 pictures)
	var userStream []models.Picture
	for _, followingID := range followingList {
		feed, err := rt.db.GetFeed(followingID)
		if err != nil {
			http.Error(w, "Failed to get feed for following", http.StatusInternalServerError)
			return
		}

		// Add up to 5 pictures from each following to the user stream
		for i := 0; i < 5 && i < len(feed); i++ {
			userStream = append(userStream, feed[i])
		}
	}

	// Send response with user stream
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userStream)
}
