package main

import (
    "fmt"
    "github.com/se-jo-ma/gox-clips/gox-clips/api/routes"
    "net/http"
)

func main() {
    router := routes.NewRouter()
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", router)
}
