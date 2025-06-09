package handler

import (
	"encoding/json"
	"net/http"
)

type WeatherResponse struct {
	City string `json:"city"`
	Temperature string `json:"temperature"`
	Condition string `json:"condition"`
	Humidity int `json:"humidity"`
	WindSpeed string `json:"windSpeed"`
	Source string `json:"source"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city");
	if city == "" {
		http.Error(w, "City is required", http.StatusBadRequest)
		return
	}

	mock := WeatherResponse{
		City: city,
		Temperature: "30C",
		Condition: "Sunny",
		Humidity: 78,
		WindSpeed: "8 km/h",
		Source: "mock",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mock)
}