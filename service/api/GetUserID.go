package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"wasaphoto/service/api/models"
)

func (rt *_router) AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var user models.User
	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Add user to the database
	UserID, err := rt.db.AddUser(user.Username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add user: %v", err), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{"message": "User added successfully", "user_id": strconv.FormatInt(UserID, 10)}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
	}
}
