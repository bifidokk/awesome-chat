package user

import (
	"context"
	"log"

	"github.com/bifidokk/awesome-chat/auth/internal/converter"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
)

// Create handles the gRPC request to create a user.
func (api *API) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create a new user: %v", req)

	userID, err := api.userService.Create(ctx, converter.ToCreateUserFromCreateRequest(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: userID,
	}, nil
}
