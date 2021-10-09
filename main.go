package main

import (
	"log"
	"net/http"

	"github.com/Nirss/fibonacci/fibonacci"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/fibonacci", getFibonacci)

	router.Run("localhost:8080")
}

type QueryParams struct {
	From int `form:"from"`
	To   int `form:"to"`
}

func getFibonacci(c *gin.Context) {
	var params QueryParams
	err := c.Bind(&params)
	if err != nil {
		log.Println("query error: ", err)
		c.String(http.StatusBadRequest, "Invalid params")
		return
	}
	fibonacciArray, err := fibonacci.FibonacciCalculation(params.From, params.To)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, fibonacciArray)
}
