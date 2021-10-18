package grpc

import (
	"context"
	"testing"

	"github.com/Nirss/fibonacci/fibonacci"

	grpcserver "github.com/Nirss/fibonacci/transport/grpc/proto"

	"github.com/Nirss/fibonacci/redis_cache"

	"github.com/stretchr/testify/assert"
)

func Test_GetRange(t *testing.T) {
	tests := []struct {
		name     string
		request  *grpcserver.GetRangeRequest
		wantBody []int32
		err      error
	}{
		{
			name: "success",
			request: &grpcserver.GetRangeRequest{
				From: 2,
				To:   10,
			},
			wantBody: []int32{1, 2, 3, 5, 8, 13, 21, 34, 55},
			err:      nil,
		},
		{
			name: "from_cannot_less_than_zero",
			request: &grpcserver.GetRangeRequest{
				From: -2,
				To:   10,
			},
			wantBody: []int32{},
			err:      fibonacci.ErrFromCannotBeLessThanZero,
		},
		{
			name: "to_cannot_be_zero",
			request: &grpcserver.GetRangeRequest{
				From: 2,
				To:   0,
			},
			wantBody: []int32{},
			err:      fibonacci.ErrToCannotBeZeroOrLess,
		},
		{
			name: "to_cannot_less_than_zero",
			request: &grpcserver.GetRangeRequest{
				From: 2,
				To:   -2,
			},
			wantBody: []int32{},
			err:      fibonacci.ErrToCannotBeZeroOrLess,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cache = redis_cache.NewCache("6379")
			service := &Service{cache: cache}
			response, err := service.GetRange(context.Background(), tt.request)
			assert.Equal(t, tt.err, err)
			if err == nil {
				assert.Equal(t, tt.wantBody, response.Result)
			}
		})
	}
}
