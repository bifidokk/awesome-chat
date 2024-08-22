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
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		loggerMock = zaptest.NewLogger(t)

		id       = gofakeit.Int64()
		password = gofakeit.Password(true, true, true, true, true, 10)
		email    = gofakeit.Email()
		name     = gofakeit.Name()

		serviceError = fmt.Errorf("service error")

		request = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            0,
		}

		response = &desc.CreateResponse{
			Id: id,
		}

		createUser = &model.CreateUser{
			Name:            name,
			Email:           email,
			Password:        password,
			ConfirmPassword: password,
			Role:            "ROLE_USER",
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateMock.Expect(ctx, createUser).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, createUser).Return(0, serviceError)
				return mock
			},
			loggerMock: loggerMock,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewUserAPI(userServiceMock, tt.loggerMock)

			result, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
