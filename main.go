package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuelkyprianou/package-calculator/calculatepacks"
)

func PacksHandler(c *gin.Context) {
	packSizes, err := readConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading config"})
		return
	}
	orderQuantityStr := c.DefaultQuery("order_quantity", "0")
	orderQuantity, err := strconv.Atoi(orderQuantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order quantity"})
		return
	}

	packsNeeded := calculatepacks.CalculatePacks(orderQuantity, packSizes)
	response := "<h1>Calculated Packs:</h1>\n"
	for size, count := range packsNeeded {
		if count > 0 {
			response += fmt.Sprintf("<p>%d x %d</p>\n", count, size)
		}
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})
	r.GET("/calculate_packs", PacksHandler)
	fmt.Println("Server started on port 8080")
	r.Run(":8080") // Run the server
}
