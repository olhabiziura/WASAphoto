package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto/service/api/models"
)

func (rt *_router) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user models.User

	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Call the database function to perform login
	userID, err := rt.db.Login(user.Username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to login user: %v", err), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "User logged in successfully",
		"user_id": strconv.FormatInt(userID, 10),
	}
	json.NewEncoder(w).Encode(response)
}
