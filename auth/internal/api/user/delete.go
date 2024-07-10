package user

import (
	"context"
	"log"

	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete handles the gRPC request to delete a user.
func (api *API) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user: %v", req)

	err := api.userService.Delete(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
