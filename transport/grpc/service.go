package grpc

import (
	"context"

	"github.com/Nirss/fibonacci/fibonacci"

	grpcserver "github.com/Nirss/fibonacci/proto"
)

type Service struct {
}

func (s *Service) GetRange(ctx context.Context, request *grpcserver.GetRangeRequest) (*grpcserver.GetRangeResponse, error) {
	result, err := fibonacci.FibonacciCalculation(int(request.From), int(request.To))
	if err != nil {
		return nil, err
	}
	var dto = make([]int32, 0, len(result))

	for _, v := range result {
		dto = append(dto, int32(v))
	}
	return &grpcserver.GetRangeResponse{Result: dto}, nil
}
