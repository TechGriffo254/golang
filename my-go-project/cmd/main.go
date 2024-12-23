package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fmt.Println("Starting the application...")

    // Define your routes and handlers here
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", nil))
}