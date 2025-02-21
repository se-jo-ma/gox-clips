package main

import (
    "fmt"
    "htmx-go-app/internal/workers"
    "time"
)

func main() {
    fmt.Println("Starting background worker...")

    for {
        workers.ProcessJobs()
        time.Sleep(5 * time.Second) // Run every 5 seconds
    }
}
