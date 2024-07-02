package converter

import (
	modelRepository "github.com/bifidokk/awesome-chat/auth/internal/repository/user/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)
import desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"

func ToUserFromRepository(user *modelRepository.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp

	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(desc.Role_value[user.Role]),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
