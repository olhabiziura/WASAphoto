package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) AddFollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	FollowerID := r.Header.Get("Authorization")
	if FollowerID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}
	
	FollowingID := ps.ByName("userID")
	if FollowingID == "" {
		http.Error(w, "Missing FollowingID parameter", http.StatusBadRequest)
		return
	}

	err := rt.db.AddFollower(FollowerID, FollowingID)
	if err != nil {
		if err.Error() == "already followed" {
			w.WriteHeader(http.StatusCreated)
			response := map[string]string{"message": "User is already being followed"}
			json.NewEncoder(w).Encode(response)
			return
		}
		http.Error(w, "Failed to follow user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
