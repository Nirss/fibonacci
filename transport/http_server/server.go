package http_server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Nirss/fibonacci/fibonacci"
	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
)

type Service struct {
	redisclient *redis.Client
}

type QueryParams struct {
	From int `form:"from"`
	To   int `form:"to"`
}

func Router(client *redis.Client) *gin.Engine {
	var redisService Service
	redisService.redisclient = client
	router := gin.Default()
	router.GET("/fibonacci", redisService.GetFibonacci)
	return router
}

func (s *Service) GetFibonacci(c *gin.Context) {
	var params QueryParams
	err := c.Bind(&params)
	if err != nil {
		log.Println("query error: ", err)
		c.String(http.StatusBadRequest, "Invalid params")
		return
	}
	redisKey := fmt.Sprintf("from=%d to=%d", params.From, params.To)
	redisValue, err := s.redisclient.Get(c, redisKey).Result()
	var fibonacciArray []int
	if err == nil {
		err = json.Unmarshal([]byte(redisValue), &fibonacciArray)
		if err != nil {
			log.Println("get redis value error: ", err)
			s.redisclient.Del(c, redisKey)
			return
		}
	} else {
		fibonacciArray, err = fibonacci.FibonacciCalculation(params.From, params.To)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		data, err := json.Marshal(fibonacciArray)
		if err != nil {
			log.Println("set redis value error: ", err)
			return
		}
		err = s.redisclient.Set(c, redisKey, string(data), 0).Err()
		if err != nil {
			log.Println("set redis value error: ", err)
		}
	}
	c.JSON(http.StatusOK, fibonacciArray)
}
