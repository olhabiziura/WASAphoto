package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto/service/api/models"
)

// Login handles user login
func (rt *_router) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user models.User

	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate username
	if user.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Call the database function to perform login
	userID, err := rt.db.Login(user.Username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to login user: %v", err), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 status code
	response := map[string]string{
		"message": "User logged in successfully",
		"user_id": strconv.FormatInt(userID, 10),
	}
	json.NewEncoder(w).Encode(response)
}
