package handlers

import (
    "encoding/json"
    "net/http"
    "example.com/mon-api/models"
)

// GET
func GetMessage(w http.ResponseWriter, r *http.Request) {
    message := models.Message{Message: "Bonjour depuis l'API!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(message)
}

// POST
func PostMessage(w http.ResponseWriter, r *http.Request) {
    var message models.Message
    err := json.NewDecoder(r.Body).Decode(&message)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    message.Message = "Message reçu: " + message.Message
    json.NewEncoder(w).Encode(message)
}
