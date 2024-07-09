package user

import (
	"context"

	"github.com/bifidokk/awesome-chat/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, user *model.UpdateUser) error {
	return s.userRepository.Update(ctx, user)
}
