package main

import (
	"github.com/Nirss/fibonacci/redis_cache"
	"github.com/Nirss/fibonacci/transport/grpc"
	"github.com/Nirss/fibonacci/transport/http_server"
)

func main() {
	var cache = redis_cache.NewCache("6379")
	router := http_server.Router(cache)
	go grpc.ListenAndServe("localhost:8081", cache)
	go router.Run("localhost:8080")
	var wait chan string
	<-wait
}
