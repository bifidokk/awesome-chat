package interceptor

import (
	"context"
	"errors"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CircuitBreakerInterceptor is an interceptor that integrates a circuit breaker to manage service availability and failure handling.
type CircuitBreakerInterceptor struct {
	cb *gobreaker.CircuitBreaker
}

// NewCircuitBreakerInterceptor creates a new CircuitBreakerInterceptor with the provided circuit breaker instance.
func NewCircuitBreakerInterceptor(cb *gobreaker.CircuitBreaker) *CircuitBreakerInterceptor {
	return &CircuitBreakerInterceptor{
		cb: cb,
	}
}

// Unary is a gRPC interceptor that uses a circuit breaker to wrap the handler execution and returns an error if the circuit is open.
func (c *CircuitBreakerInterceptor) Unary(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := c.cb.Execute(func() (interface{}, error) {
		return handler(ctx, req)
	})

	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return nil, status.Error(codes.Unavailable, "service unavailable")
		}

		return nil, err
	}

	return res, nil
}
