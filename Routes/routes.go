package routes

import (
    "github.com/gorilla/mux"
    "example.com/mon-api/handlers"
)

func InitRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/api/message", handlers.GetMessage).Methods("GET")
    router.HandleFunc("/api/message", handlers.PostMessage).Methods("POST")

    return router
}
