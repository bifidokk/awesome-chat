package user

import (
	"context"

	"github.com/bifidokk/awesome-chat/auth/internal/converter"
	"github.com/bifidokk/awesome-chat/auth/internal/logger"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"go.uber.org/zap"
)

// Get handles the gRPC request to get a user.
func (api *API) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	logger.Info("Get user", zap.Any("request", req))

	user, err := api.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return converter.ToGetUserResponseFromUser(user), nil
}
