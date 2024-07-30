package chat

import (
	"context"
	"log"

	desc "github.com/bifidokk/awesome-chat/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete handles the gRPC request to delete a chat.
func (api *API) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete a chat: %v", req)

	err := api.chatService.Delete(ctx, req.Id)

	return &emptypb.Empty{}, err
}
