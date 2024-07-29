package chat

import (
	"github.com/bifidokk/awesome-chat/chat-server/internal/repository"
	"github.com/bifidokk/awesome-chat/chat-server/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
}

// NewChatService creates a new instance of ChatService.
func NewChatService(chatRepository repository.ChatRepository) service.ChatService {
	return &serv{chatRepository}
}
