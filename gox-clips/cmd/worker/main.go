package main

import (
    "fmt"
    "github.com/se-jo-ma/gox-clips/gox-clips/internal/workers"
    "time"
)

func main() {
    fmt.Println("Starting background worker...")

    for {
        workers.ProcessJobs()
        time.Sleep(5 * time.Second) // Run every 5 seconds
    }
}
