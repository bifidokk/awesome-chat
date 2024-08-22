package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/bifidokk/awesome-chat/auth/internal/api/user"
	"github.com/bifidokk/awesome-chat/auth/internal/model"
	"github.com/bifidokk/awesome-chat/auth/internal/service"
	serviceMocks "github.com/bifidokk/awesome-chat/auth/internal/service/mocks"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		loggerMock = zaptest.NewLogger(t)

		id    = gofakeit.Int64()
		email = gofakeit.Email()
		name  = gofakeit.Name()

		serviceError = fmt.Errorf("service error")

		request = &desc.UpdateRequest{
			Id: id,
			Name: &wrapperspb.StringValue{
				Value: name,
			},
			Email: &wrapperspb.StringValue{
				Value: email,
			},
		}

		response = &emptypb.Empty{}

		updateUser = &model.UpdateUser{
			ID:    id,
			Name:  name,
			Email: email,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock userServiceMockFunc
		loggerMock      *zap.Logger
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: response,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, updateUser).Return(nil)
				return mock
			},
			loggerMock: loggerMock,
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: nil,
			err:  serviceError,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, updateUser).Return(serviceError)
				return mock
			},
			loggerMock: loggerMock,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewUserAPI(userServiceMock, loggerMock)

			result, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
