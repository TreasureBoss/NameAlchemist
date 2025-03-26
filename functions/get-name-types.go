package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Config struct {
	NameTypes []string `json:"nameTypes"`
}

func loadConfig() (Config, error) {
	file, err := os.ReadFile("./config/name-config.json")
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	return config, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	config, err := loadConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(config.NameTypes)
}

func main() {
	http.HandleFunc("/", handler)
}