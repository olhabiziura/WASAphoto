package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (rt *_router) DeleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	UserID := r.Header.Get("Authorization")
	if UserID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	BannedID := ps.ByName("userID")
	if BannedID == "" {
		http.Error(w, "Missing BannedID parameter", http.StatusBadRequest)
		return
	}

	err := rt.db.DeleteBan(UserID, BannedID)
	if err != nil {
		http.Error(w, "Failed to unban user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "User unbanned successfully"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
	}
}
