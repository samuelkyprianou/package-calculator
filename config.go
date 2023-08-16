package main

import (
	"encoding/json"
	"os"
)

func readConfig() (map[int]int, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config struct {
		PackSizes []int `json:"packSizes"`
	}
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	order := make(map[int]int)
	for _, size := range config.PackSizes {
		order[size] = 0
	}
	return order, nil
}