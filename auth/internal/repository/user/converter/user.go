package converter

import (
	"github.com/bifidokk/awesome-chat/auth/internal/model"
	modelRepository "github.com/bifidokk/awesome-chat/auth/internal/repository/user/model"
)

// ToUserFromRepository converts a User model from the repository layer to a User model for the business logic layer.
func ToUserFromRepository(user *modelRepository.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
