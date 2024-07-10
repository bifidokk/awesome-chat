package user

import (
	"context"
	"log"

	"github.com/bifidokk/awesome-chat/auth/internal/converter"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
)

// Get handles the gRPC request to get a user.
func (api *API) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get user: %v", req)

	user, err := api.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return converter.ToGetUserResponseFromUser(user), nil
}
