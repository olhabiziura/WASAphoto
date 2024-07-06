package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto/service/api/models"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func (rt *_router) GetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var profile models.Profile

	// Get UserID from Authorization header
	UserID := r.Header.Get("Authorization")
	if UserID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Get ProfileID from URL parameter
	ProfileID := ps.ByName("userID")
	if ProfileID == "" {
		http.Error(w, "Missing userID parameter", http.StatusBadRequest)
		return
	}

	// Set Profile UserID
	profile.UserID = ProfileID

	// Get username for the profile
	username, err := rt.db.GetUsername(ProfileID)
	if err != nil {
		http.Error(w, "Failed to retrieve username", http.StatusInternalServerError)
		return
	}
	profile.Username = username

	// Get followers for the profile
	followerList, err := rt.db.GetFollowers(ProfileID)
	if err != nil {
		http.Error(w, "Failed to retrieve followers", http.StatusInternalServerError)
		return
	}
	profile.FollowerList = followerList

	// Get following list for the profile
	followingList, err := rt.db.GetFollowing(ProfileID)
	if err != nil {
		http.Error(w, "Failed to retrieve following", http.StatusInternalServerError)
		return
	}
	profile.FollowingList = followingList

	// Only get ban list if the profile is the same as the logged-in user
	if UserID == ProfileID {
		banList, err := rt.db.GetBan(ProfileID)
		fmt.Println(err)
		if err != nil {
			http.Error(w, "Failed to retrieve ban list", http.StatusInternalServerError)
			return
		}
		profile.BanList = banList
	}

	// Get feed for the profile
	feed, err := rt.db.GetFeed(ProfileID)
	if err != nil {
		http.Error(w, "Failed to retrieve feed", http.StatusInternalServerError)
		return
	}
	

	for i, picture := range feed {

		encodedImage, err := ReadImageAsBase64(picture.Address)
		if err != nil {
			http.Error(w, "Failed to read image", http.StatusInternalServerError)
			return
		}
		feed[i].Image = encodedImage
	}

	profile.PhotoList = feed

	// Encode profile instance in the response and send it back
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)

	
}
