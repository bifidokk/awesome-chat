package chat

import (
	"context"
	"log"

	"github.com/bifidokk/awesome-chat/chat-server/internal/converter"
	desc "github.com/bifidokk/awesome-chat/chat-server/pkg/chat_v1"
)

// Create handles the gRPC request to create a chat.
func (api *API) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create a new chat: %v", req)

	chatID, err := api.chatService.Create(ctx, converter.ToCreateChatFromCreateRequest(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: chatID,
	}, nil
}
