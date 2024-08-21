package user

import (
	"context"

	"github.com/bifidokk/awesome-chat/auth/internal/logger"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete handles the gRPC request to delete a user.
func (api *API) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	logger.Info("Delete user", zap.Any("request", req))

	err := api.userService.Delete(ctx, req.Id)

	return &emptypb.Empty{}, err
}
