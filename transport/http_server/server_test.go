package http_server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nirss/fibonacci/redis_cache"

	"github.com/stretchr/testify/assert"
)

func Test_Router(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		wantBody   string
		wantStatus int
	}{
		{
			name:       "success",
			url:        "/fibonacci?from=2&to=6",
			wantBody:   "[1,2,3,5,8]",
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid_from_fibonacci_params",
			url:        "/fibonacci?from=-2&to=6",
			wantBody:   "numbers from cannot be less than zero",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "to_cannot_be_zero",
			url:        "/fibonacci?from=2&to=0",
			wantBody:   "numbers to cannot be zero or less",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "to_cannot_less_than_zero",
			url:        "/fibonacci?from=2&to=-1",
			wantBody:   "numbers to cannot be zero or less",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid_query_params",
			url:        "/fibonacci?from=Hello&to=6",
			wantBody:   "Invalid params",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cache = redis_cache.NewCache("6379")
			router := Router(cache)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.url, nil)
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}
