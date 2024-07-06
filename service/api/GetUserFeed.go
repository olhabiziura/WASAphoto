package api

import (
	"encoding/json"
	"net/http"
	"encoding/base64"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	//"path/filepath"
	"fmt"


)





// ReadImageAsBase64 reads an image from the file path and returns it as a base64 encoded string
func ReadImageAsBase64(filePath string) (string, error) {
	imageData, err := ioutil.ReadFile(filePath)
	fmt.Println(err)
	if err != nil {
		
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imageData), nil
}


func (rt *_router) GetUserFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get UserID from Authorization header
	userID := r.Header.Get("Authorization")
	if userID == "" {
		http.Error(w, "Missing authorization header", http.StatusBadRequest)
		return
	}

	// Get feed for the user
	feed, err := rt.db.GetFeed(userID)
	if err != nil {
		switch err.Error() {
		case "invalid user_id":
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
		default:
			http.Error(w, "Failed to get feed", http.StatusInternalServerError)
		}
		
	}
	fmt.Println(feed)
	// Fetch and encode pictures
	for i, picture := range feed {

		fmt.Println("kjnljn")
		encodedImage, err := ReadImageAsBase64(picture.Address)
		if err != nil {
			http.Error(w, "Failed to read image", http.StatusInternalServerError)
			return
		}
		feed[i].Image = encodedImage
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Marshal feed to JSON and send response
	json.NewEncoder(w).Encode(feed)
}