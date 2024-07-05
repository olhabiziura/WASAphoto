package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) AddPhotoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get UserID from Authorization header
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Parse multipart form to get file and description
	err := r.ParseMultipartForm(10 << 20) // maxMemory 10MB
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse multipart form: %v", err), http.StatusBadRequest)
		return
	}

	// Get file from form
	file, handler, err := r.FormFile("picture")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get file from form: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate unique filename using timestamp and user ID
	fileName := fmt.Sprintf("%s_%d%s", userID, time.Now().UnixNano(), filepath.Ext(handler.Filename))
	filePath := filepath.Join("cmd/photos", fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create file: %v", err), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save file: %v", err), http.StatusInternalServerError)
		return
	}

	// Get description from form
	description := r.FormValue("description")

	// Call database function to add photo record
	err = rt.db.AddPhoto(userID, filePath, time.Now(), description)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add photo record: %v", err), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "The picture was successfully added to the user's feed.",
		"file":    filePath, // Provide some reference to the uploaded file if needed
	}
	json.NewEncoder(w).Encode(response)
}
