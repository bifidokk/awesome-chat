package service

import (
	"context"
	"github.com/bifidokk/awesome-chat/auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, user *model.CreateUser) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.UpdateUser) error
	Delete(ctx context.Context, id int64) error
}
