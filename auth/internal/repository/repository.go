package repository

import (
	"context"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
)

type UserRepository interface {
	Create(ctx context.Context, data *desc.CreateRequest) (int64, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, data *desc.UpdateRequest) error
	Get(ctx context.Context, id int64) (*desc.GetResponse, error)
}
