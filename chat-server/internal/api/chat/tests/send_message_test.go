package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/bifidokk/awesome-chat/chat-server/internal/api/chat"
	"github.com/bifidokk/awesome-chat/chat-server/internal/model"
	"github.com/bifidokk/awesome-chat/chat-server/internal/service"
	serviceMocks "github.com/bifidokk/awesome-chat/chat-server/internal/service/mocks"
	desc "github.com/bifidokk/awesome-chat/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id           = gofakeit.Int64()
		date         = gofakeit.Date()
		text         = gofakeit.Comment()
		fromName     = gofakeit.Name()
		serviceError = fmt.Errorf("service error")

		request = &desc.SendMessageRequest{
			From: fromName,
			Text: text,
			Timestamp: &timestamppb.Timestamp{
				Seconds: date.Unix(),
				Nanos:   int32(date.Nanosecond()),
			},
			ChatID: id,
		}

		response = &emptypb.Empty{}

		sendMessage = &model.SendMessage{
			From:   fromName,
			Text:   text,
			ChatID: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: response,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, sendMessage).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: response,
			err:  serviceError,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, sendMessage).Return(serviceError)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewChatAPI(chatServiceMock)

			result, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
