package service

import (
	"context"

	"github.com/bifidokk/awesome-chat/chat-server/internal/model"
)

// ChatService defines the methods that any service providing chat-related operations should implement.
type ChatService interface {
	Create(ctx context.Context, data *model.CreateChat) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, id int64, message *model.SendMessage) error
}
