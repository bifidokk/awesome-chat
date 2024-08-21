package user

import (
	"context"

	"github.com/bifidokk/awesome-chat/auth/internal/converter"
	"github.com/bifidokk/awesome-chat/auth/internal/logger"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update handles the gRPC request to update a user.
func (api *API) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	logger.Info("Update user", zap.Any("request", req))

	err := api.userService.Update(ctx, converter.ToUpdateUserFromUpdateRequest(req))

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
