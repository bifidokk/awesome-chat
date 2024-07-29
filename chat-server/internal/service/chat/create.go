package chat

import (
	"context"

	"github.com/bifidokk/awesome-chat/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, chat *model.CreateChat) (int64, error) {
	chatID, err := s.chatRepository.Create(ctx, chat)

	if err != nil {
		return 0, err
	}

	return chatID, nil
}
