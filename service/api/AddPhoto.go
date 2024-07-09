package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"wasaphoto/service/api/models"
	"log"
)

func (rt *_router) AddPhotoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var picture models.Picture

	// Get UserID from Authorization header
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Parse multipart form to get file and description
	err := r.ParseMultipartForm(10 << 20) // maxMemory 10MB
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse multipart form: %w", err), http.StatusBadRequest)
		return
	}

	// Get file from form
	file, handler, err := r.FormFile("picture")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get file from form: %w", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate unique filename using timestamp and user ID
	fileName := fmt.Sprintf("%s_%d%s", userID, time.Now().UnixNano(), filepath.Ext(handler.Filename))
	filePath := filepath.Join("cmd/photos", fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create file: %w", err), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save file: %w", err), http.StatusInternalServerError)
		return
	}

	// Get description from form
	description := r.FormValue("description")

	// Call database function to add photo record
	pictureID, err := rt.db.AddPhoto(userID, filePath, time.Now(), description)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add photo record: %w", err), http.StatusInternalServerError)
		return
	}

	// Populate picture struct
	picture.OwnerID = userID
	picture.Date = time.Now()
	picture.PictureID = strconv.FormatInt(pictureID, 10)
	picture.Address = filePath

	var username string
	username, err = rt.db.GetUsername(userID)
	picture.Username = username

	encodedImage, err := ReadImageAsBase64(picture.Address)
	if err != nil {
		http.Error(w, "Failed to read image", http.StatusInternalServerError)
		return
	}

	picture.Image = encodedImage

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(picture)
	if err != nil {
		http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %w", err)
	}
}
