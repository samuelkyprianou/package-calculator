package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func PacksHandler(w http.ResponseWriter, r *http.Request) {
	packSizes, err := readConfig()
	if err != nil {
		http.Error(w, "Error reading config", http.StatusInternalServerError)
		return
	}
	query := r.URL.Query()
	orderQuantityStr := query.Get("order_quantity")
	orderQuantity, err := strconv.Atoi(orderQuantityStr)
	if err != nil {
		http.Error(w, "Invalid order quantity", http.StatusBadRequest)
		return
	}
	
	packsNeeded := CalculatePacks(orderQuantity, packSizes)
	response := "<h1>Calculated Packs:</h1>\n"
	for size, count := range packsNeeded {
		if count > 0 {
			response += fmt.Sprintf("<p>%d x %d</p>\n", count, size)
		}
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/calculate_packs", PacksHandler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

