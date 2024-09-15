package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rateLimiter "github.com/bifidokk/awesome-chat/auth/internal/rate_limiter"
)

// RateLimiterInterceptor is an interceptor that applies rate limiting using a token bucket algorithm.
type RateLimiterInterceptor struct {
	rateLimiter *rateLimiter.TokenBucketLimiter
}

// NewRateLimiterInterceptor creates a new instance of RateLimiterInterceptor with the provided token bucket rate limiter.
func NewRateLimiterInterceptor(rateLimiter *rateLimiter.TokenBucketLimiter) *RateLimiterInterceptor {
	return &RateLimiterInterceptor{rateLimiter: rateLimiter}
}

// Allow is a gRPC interceptor method that checks if the rate limiter allows processing the request; otherwise, it returns a "ResourceExhausted" error.
func (r *RateLimiterInterceptor) Allow(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !r.rateLimiter.Allow() {
		return nil, status.Error(codes.ResourceExhausted, "too many requests")
	}

	return handler(ctx, req)
}
