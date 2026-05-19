package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	UPC      int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var products = []product{
	{UPC: 2590, Name: "Vizio 40 inch HDTV", Price: 48.99, Quantity: 20},
	{UPC: 2386, Name: "TCL Roku50 inch Smart TV", Price: 75.99, Quantity: 50},
	{UPC: 3456, Name: "MS Bath Mat Blue", Price: 5.99, Quantity: 10},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.Run("localhost:8080")
	log.Println("Server started at http://localhost:8080")
}
