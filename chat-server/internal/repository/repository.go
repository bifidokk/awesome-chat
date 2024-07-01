package repository

import (
	"context"
	desc "github.com/bifidokk/awesome-chat/chat-server/pkg/chat_v1"
)

type ChatRepository interface {
	Create(ctx context.Context, data *desc.CreateRequest) (int64, error)
	Delete(ctx context.Context, id int64) error
}
