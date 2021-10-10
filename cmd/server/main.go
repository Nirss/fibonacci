package main

import (
	"github.com/Nirss/fibonacci/transport/grpc"
	"github.com/Nirss/fibonacci/transport/http_server"
	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	router := http_server.Router(client)
	go grpc.ListenAndServe("localhost:8081", client)
	go router.Run("localhost:8080")
	var wait chan string
	<-wait
}
