package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) AddBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	UserID := r.Header.Get("Authorization")
	if UserID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	BanID := ps.ByName("userID")
	if BanID == "" {
		http.Error(w, "Missing BanID parameter", http.StatusBadRequest)
		return
	}

	err := rt.db.AddBan(UserID, BanID)
	if err != nil {
		if err.Error() == "already banned" {
			w.WriteHeader(http.StatusCreated)
			response := map[string]string{"message": "User is already banned"}
			json.NewEncoder(w).Encode(response)
			return
		}
		http.Error(w, "Failed to ban user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
