package api

import (
    "encoding/json"
    "net/http"
    "github.com/julienschmidt/httprouter"

)

func (rt *_router) DeleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        
    // Extract UserID from request headers
	UserID := r.Header.Get("Authorization")
	if UserID == "" {
		http.Error(w, "Missing user_id in headers", http.StatusBadRequest)
		return
	}

    // Extract PictureID from URL parameters
    PictureID := ps.ByName("pictureID")
	if PictureID == "" {
		http.Error(w, "Missing pictureID parameter", http.StatusBadRequest)
		return
	}


    err := rt.db.DeleteLike(UserID, PictureID)
    if err != nil {
		http.Error(w, "Failed to delete like", http.StatusInternalServerError)
		return
	}

    // Send success response
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    response := map[string]string{"message": "Like removed successfully"}
    json.NewEncoder(w).Encode(response)
}
