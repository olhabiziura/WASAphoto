package api

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func (rt *_router) AddLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extract UserID from request headers
    userID := r.Header.Get("Authorization")
    if userID == "" {
        http.Error(w, "Missing authorization header", http.StatusBadRequest)
        return
    }

    // Extract PictureID from URL parameters
    pictureID := ps.ByName("PictureID")
    if pictureID == "" {
        http.Error(w, "Missing pictureID parameter", http.StatusBadRequest)
        return
    }

    // Insert into the database
    err := rt.db.AddLike(userID, pictureID)
    if err != nil {
        http.Error(w, "Failed to add like to a picture", http.StatusInternalServerError)
        return
    }

    // Send success response
    w.WriteHeader(http.StatusNoContent)
}
