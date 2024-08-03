package user

import (
	"github.com/bifidokk/awesome-chat/auth/internal/service"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
)

// API is a struct that implements the AuthV1Server gRPC interface.
type API struct {
	desc.UnimplementedAuthV1Server
	userService service.UserService
}

// NewUserAPI creates a new instance of UserAPI.
func NewUserAPI(userService service.UserService) *API {
	return &API{userService: userService}
}
