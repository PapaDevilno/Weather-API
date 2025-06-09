package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"weather-api/handler"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	http.HandleFunc("/weather", handler.GetWeather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}