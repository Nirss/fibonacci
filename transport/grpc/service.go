package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/Nirss/fibonacci/fibonacci"

	"github.com/go-redis/redis/v8"

	"google.golang.org/grpc"

	grpcserver "github.com/Nirss/fibonacci/transport/grpc/proto"
)

var (
	ErrUnexpectedError = errors.New("unexpected error, please repeat again later")
)

type Service struct {
	grpcserver.UnimplementedFibonacciServiceServer
	redisclient *redis.Client
}

func (s *Service) GetRange(ctx context.Context, request *grpcserver.GetRangeRequest) (*grpcserver.GetRangeResponse, error) {
	redisKey := fmt.Sprintf("from=%d to=%d", request.From, request.To)
	redisValue, err := s.redisclient.Get(ctx, redisKey).Result()
	var result []int
	if err == nil {
		err = json.Unmarshal([]byte(redisValue), &result)
		if err != nil {
			log.Println("get redis value error: ", err)
			s.redisclient.Del(ctx, redisKey)
			return nil, ErrUnexpectedError
		}
	} else {
		result, err = fibonacci.FibonacciCalculation(int(request.From), int(request.To))
		if err != nil {
			return nil, err
		}
		data, err := json.Marshal(result)
		if err != nil {
			log.Println("set redis value error: ", err)
			return nil, err
		}
		err = s.redisclient.Set(ctx, redisKey, string(data), 0).Err()
		if err != nil {
			log.Println("set redis value error: ", err)
		}
	}
	var dto = make([]int32, 0, len(result))

	for _, v := range result {
		dto = append(dto, int32(v))
	}
	return &grpcserver.GetRangeResponse{Result: dto}, nil
}

func ListenAndServe(port string, client *redis.Client) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	opt := []grpc.ServerOption{}
	server := grpc.NewServer(opt...)
	grpcserver.RegisterFibonacciServiceServer(server, &Service{redisclient: client})
	server.Serve(listener)
}
