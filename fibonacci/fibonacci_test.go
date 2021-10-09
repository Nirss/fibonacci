package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciCalculation(t *testing.T) {
	tests := []struct {
		name    string
		From    int
		To      int
		want    []int
		wantErr error
	}{
		{
			name:    "success",
			From:    1,
			To:      8,
			want:    []int{0, 1, 1, 2, 3, 5, 8, 13},
			wantErr: nil,
		},
		{
			name:    "from_greater_than_to",
			From:    2,
			To:      1,
			want:    []int{},
			wantErr: ErrFromGreaterThanTo,
		},
		{
			name:    "from_cannot_be_zero",
			From:    0,
			To:      1,
			want:    []int{},
			wantErr: ErrFromOrToCannotBeZeroOrLess,
		},
		{
			name:    "to_cannot_be_zero",
			From:    1,
			To:      0,
			want:    []int{},
			wantErr: ErrFromOrToCannotBeZeroOrLess,
		},
		{
			name:    "from_cannot_less_than_zero",
			From:    -1,
			To:      2,
			want:    []int{},
			wantErr: ErrFromOrToCannotBeZeroOrLess,
		},
		{
			name:    "to_cannot_less_than_zero",
			From:    2,
			To:      -2,
			want:    []int{},
			wantErr: ErrFromOrToCannotBeZeroOrLess,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RecFibonacciCalculation(tt.From, tt.To)
			assert.EqualValues(t, tt.wantErr, err)
			assert.EqualValues(t, tt.want, result)
		})
	}
}

func BenchmarkFibonacciCalculation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecFibonacciCalculation(1, 25)
	}
}
