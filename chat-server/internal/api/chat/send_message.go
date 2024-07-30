package chat

import (
	"context"

	"github.com/bifidokk/awesome-chat/chat-server/internal/converter"
	desc "github.com/bifidokk/awesome-chat/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage handles the gRPC request to send a message in a chat.
func (api *API) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := api.chatService.SendMessage(
		ctx,
		0,
		converter.ToSendMessageFromSendMessageRequest(req),
	)

	return &emptypb.Empty{}, err
}
