package chat

import (
	"context"
	"log"
	"strconv"

	"github.com/bifidokk/awesome-chat/chat-server/internal/model"
)

func (s *serv) SendMessage(_ context.Context, data *model.SendMessage) error {
	log.Println("sending message " + data.Text + " from " + data.From + " to chat" + strconv.FormatInt(data.ChatID, 10))

	return nil
}
