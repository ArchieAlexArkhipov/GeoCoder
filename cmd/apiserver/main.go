package main

import (
    "log"
    "github.com/joho/godotenv"
    "geocoder/internal/apiserver"
)

// init is invoked before main()
func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
    server := apiserver.APIServerFactory()

    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}
