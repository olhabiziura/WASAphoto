package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
		"log"
)

func (rt *_router) SearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract username from query parameters
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Missing username parameter", http.StatusBadRequest)
		return
	}

	users, err := rt.db.GetUsers(username)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
	}
}
