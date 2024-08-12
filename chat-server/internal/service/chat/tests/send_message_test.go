package tests

import (
	"context"
	"testing"

	"github.com/bifidokk/awesome-chat/chat-server/internal/model"
	"github.com/bifidokk/awesome-chat/chat-server/internal/repository"
	repositoryMocks "github.com/bifidokk/awesome-chat/chat-server/internal/repository/mocks"
	chatService "github.com/bifidokk/awesome-chat/chat-server/internal/service/chat"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository

	type args struct {
		ctx         context.Context
		sendMessage *model.SendMessage
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		sendMessage = &model.SendMessage{
			From:   gofakeit.Name(),
			Text:   gofakeit.Comment(),
			ChatID: id,
		}
	)

	tests := []struct {
		name               string
		args               args
		err                error
		chatRepositoryMock chatRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:         ctx,
				sendMessage: sendMessage,
			},
			err: nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			srv := chatService.NewChatService(chatRepositoryMock)

			err := srv.SendMessage(tt.args.ctx, tt.args.sendMessage)
			require.Equal(t, tt.err, err)
		})
	}
}
