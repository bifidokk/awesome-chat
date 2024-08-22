package user

import (
	"github.com/bifidokk/awesome-chat/auth/internal/service"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"go.uber.org/zap"
)

// API is a struct that implements the AuthV1Server gRPC interface.
type API struct {
	desc.UnimplementedAuthV1Server
	userService service.UserService
	logger      *zap.Logger
}

// NewUserAPI creates a new instance of UserAPI.
func NewUserAPI(userService service.UserService, logger *zap.Logger) *API {
	return &API{userService: userService, logger: logger}
}
