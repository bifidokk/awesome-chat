package repository

import (
	"context"

	"github.com/bifidokk/awesome-chat/chat-server/internal/model"
)

// ChatRepository defines the methods that any repository handling chat data storage should implement.
type ChatRepository interface {
	Create(ctx context.Context, data *model.CreateChat) (int64, error)
	Delete(ctx context.Context, id int64) error
}
