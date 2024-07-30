package chat

import (
	"github.com/bifidokk/awesome-chat/chat-server/internal/service"
	desc "github.com/bifidokk/awesome-chat/chat-server/pkg/chat_v1"
)

// API is a struct that implements the ChatV1Server gRPC interface.
type API struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewChatAPI creates a new instance of ChatApi.
func NewChatAPI(chatService service.ChatService) *API {
	return &API{chatService: chatService}
}
