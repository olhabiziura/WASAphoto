package api

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) DeleteFollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	err := rt.db.DeleteFollower(FollowerID, FollowingID)
	if err != nil {
		if err.Error() == "not following" {
			w.WriteHeader(http.StatusCreated)
			response := map[string]string{"message": "User wasn't following the provided user"}
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
				log.Printf("Failed to encode response: %w", err)
			}
			return
		}
		http.Error(w, "Failed to unfollow user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
