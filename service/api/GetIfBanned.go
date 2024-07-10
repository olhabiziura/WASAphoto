package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (rt *_router) GetIfBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get UserID from Authorization header
	UserID := r.Header.Get("Authorization")
	if UserID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Get BannerID from URL parameter
	BannerID := ps.ByName("userID")
	if BannerID == "" {
		http.Error(w, "Missing userID parameter", http.StatusBadRequest)
		return
	}

	// Check if the user is banned
	IsBanned, err := rt.db.GetIfBan(UserID, BannerID)
	if err != nil {
		http.Error(w, "Failed to check ban status", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]bool{"isBanned": IsBanned})
	if err != nil {
		http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
	}
}
