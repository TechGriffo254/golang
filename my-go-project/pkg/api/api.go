package api

import (
    "net/http"
)

// Handler function for the root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to my Go project API!"))
}

// InitializeRoutes sets up the API routes
func InitializeRoutes() {
    http.HandleFunc("/", RootHandler)
}