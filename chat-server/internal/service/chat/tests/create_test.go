package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/bifidokk/awesome-chat/chat-server/internal/model"
	"github.com/bifidokk/awesome-chat/chat-server/internal/repository"
	repositoryMocks "github.com/bifidokk/awesome-chat/chat-server/internal/repository/mocks"
	chatService "github.com/bifidokk/awesome-chat/chat-server/internal/service/chat"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository

	type args struct {
		ctx        context.Context
		createChat *model.CreateChat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		repositoryError = fmt.Errorf("repository error")

		id    = gofakeit.Int64()
		names = []string{
			gofakeit.Name(),
			gofakeit.Name(),
		}

		createChat = &model.CreateChat{
			Usernames: names,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		chatRepositoryMock chatRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:        ctx,
				createChat: createChat,
			},
			want: id,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, createChat).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx:        ctx,
				createChat: createChat,
			},
			want: 0,
			err:  repositoryError,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, createChat).Return(0, repositoryError)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			srv := chatService.NewChatService(chatRepositoryMock)

			result, err := srv.Create(tt.args.ctx, tt.args.createChat)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
