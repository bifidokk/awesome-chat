package user

import (
	"github.com/bifidokk/awesome-chat/auth/internal/repository"
	"github.com/bifidokk/awesome-chat/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &serv{userRepository}
}
