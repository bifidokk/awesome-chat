package user

import (
	"context"

	"github.com/bifidokk/awesome-chat/auth/internal/converter"
	"github.com/bifidokk/awesome-chat/auth/internal/logger"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"go.uber.org/zap"
)

// Create handles the gRPC request to create a user.
func (api *API) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	logger.Info("Create a new user", zap.Any("request", req))

	userID, err := api.userService.Create(ctx, converter.ToCreateUserFromCreateRequest(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: userID,
	}, nil
}
