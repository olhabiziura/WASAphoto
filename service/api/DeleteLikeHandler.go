package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) DeleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract UserID from request headers
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Extract PictureID from URL parameters
	pictureID := ps.ByName("pictureID")
	if pictureID == "" {
		http.Error(w, "Missing pictureID parameter", http.StatusBadRequest)
		return
	}

	// Delete the like from the database
	err := rt.db.DeleteLike(userID, pictureID)
	if err != nil {
		http.Error(w, "Failed to delete like", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusNoContent)
}
