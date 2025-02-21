package main

import (
    "fmt"
    "htmx-go-app/api/routes"
    "net/http"
)

func main() {
    router := routes.NewRouter()
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", router)
}
