package converter

import (
	"github.com/bifidokk/awesome-chat/chat-server/internal/model"
	desc "github.com/bifidokk/awesome-chat/chat-server/pkg/chat_v1"
)

// ToCreateChatFromCreateRequest converts a CreateRequest from the gRPC layer to a CreateChat model for the business logic layer
func ToCreateChatFromCreateRequest(req *desc.CreateRequest) *model.CreateChat {
	return &model.CreateChat{
		Usernames: req.Usernames,
	}
}

// ToSendMessageFromSendMessageRequest converts a SendMessageRequest from the gRPC layer to a SendMessage model for the business logic layer.
func ToSendMessageFromSendMessageRequest(req *desc.SendMessageRequest) *model.SendMessage {
	return &model.SendMessage{
		Text:   req.Text,
		From:   req.From,
		ChatID: req.ChatID,
	}
}
