package http_server

import (
	"log"
	"net/http"

	"github.com/Nirss/fibonacci/redis_cache"

	"github.com/Nirss/fibonacci/fibonacci"
	"github.com/gin-gonic/gin"
)

type FibonacciServer struct {
	cache *redis_cache.Cache
}

type QueryParams struct {
	From int `form:"from"`
	To   int `form:"to"`
}

func Router(cache *redis_cache.Cache) *gin.Engine {
	router := gin.Default()
	var server FibonacciServer
	server.cache = cache
	router.GET("/fibonacci", server.GetFibonacci)
	return router
}

func (f *FibonacciServer) GetFibonacci(c *gin.Context) {
	var params QueryParams
	err := c.Bind(&params)
	if err != nil {
		log.Println("query error: ", err)
		c.String(http.StatusBadRequest, "Invalid params")
		return
	}
	result, err := f.cache.GetValue(c, params.From, params.To)
	if err != nil {
		result, err = fibonacci.FibonacciCalculation(params.From, params.To)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = f.cache.SetValue(c, params.From, params.To, result)
		if err != nil {
			log.Println("set redis value error: ", err)
		}
	}
	c.JSON(http.StatusOK, result)
}
