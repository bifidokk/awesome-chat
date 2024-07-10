package user

import (
	"context"
	"log"

	"github.com/bifidokk/awesome-chat/auth/internal/converter"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update handles the gRPC request to update a user.
func (api *API) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update user: %v", req)

	err := api.userService.Update(ctx, converter.ToUpdateUserFromUpdateRequest(req))

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
