package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

// ValidateInterceptor is a gRPC interceptor that validates incoming requests.
func ValidateInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if val, ok := req.(validator); ok {
		if err := val.Validate(); err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
