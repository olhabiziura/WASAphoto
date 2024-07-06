package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wasaphoto/service/api/models"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

func (rt *_router) AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        var UserID int64
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
		
		fmt.Println("adjsn;ojn")
		// Send success response
		w.Header().Set("Content-Type", "application/json")
		
		response := map[string]string{"message": "User added successfully", "user_id": strconv.FormatInt(UserID,10)}
		json.NewEncoder(w).Encode(response)
	}

