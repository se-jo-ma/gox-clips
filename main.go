package main

import (
    "log"

    "gox-clips/internal/config"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Use the configuration
    log.Printf("Database Host: %s", cfg.DBHost)
}
