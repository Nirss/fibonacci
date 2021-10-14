package grpc

import (
	"context"
	"log"
	"net"

	"github.com/Nirss/fibonacci/redis_cache"

	"github.com/Nirss/fibonacci/fibonacci"

	"google.golang.org/grpc"

	grpcserver "github.com/Nirss/fibonacci/transport/grpc/proto"
)

type Service struct {
	grpcserver.UnimplementedFibonacciServiceServer
	cache *redis_cache.Cache
}

func (s *Service) GetRange(ctx context.Context, request *grpcserver.GetRangeRequest) (*grpcserver.GetRangeResponse, error) {
	result, err := s.cache.GetValue(ctx, int(request.From), int(request.To))
	if err != nil {
		result, err = fibonacci.FibonacciCalculation(int(request.From), int(request.To))
		if err != nil {
			return nil, err
		}
		err = s.cache.SetValue(ctx, int(request.From), int(request.To), result)
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

func ListenAndServe(port string, cache *redis_cache.Cache) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	opt := []grpc.ServerOption{}
	server := grpc.NewServer(opt...)
	grpcserver.RegisterFibonacciServiceServer(server, &Service{cache: cache})
	server.Serve(listener)
}
