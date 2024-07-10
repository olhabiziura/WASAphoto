package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) DeletePhotoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get UserID from Authorization header
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}

	// Get photoID from URL parameters
	photoID := ps.ByName("pictureID")
	if photoID == "" {
		http.Error(w, "Missing pictureID parameter", http.StatusBadRequest)
		return
	}

	// Call database function to delete photo record and get file path
	filePath, err := rt.db.DeletePhoto(userID, photoID)
	if err != nil {
		if err.Error() == "photo not found" {
			http.Error(w, "Photo not found", http.StatusBadRequest)
		} else {
			http.Error(w, fmt.Sprintf("Failed to delete photo: %v", err), http.StatusInternalServerError)
		}
		return
	}

	// Delete the file from disk
	err = os.Remove(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete file: %v", err), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusNoContent)
}
