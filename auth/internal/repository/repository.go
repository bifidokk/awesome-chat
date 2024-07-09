package repository

import (
	"context"
	"github.com/bifidokk/awesome-chat/auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, data *model.CreateUser) (int64, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, data *model.UpdateUser) error
	Get(ctx context.Context, id int64) (*model.User, error)
}
