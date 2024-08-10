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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		email = gofakeit.Email()
		name  = gofakeit.Name()

		serviceError = fmt.Errorf("service error")

		request = &desc.GetRequest{
			Id: id,
		}

		date = gofakeit.Date()

		response = &desc.GetResponse{
			Id:    id,
			Name:  name,
			Email: email,
			Role:  desc.Role(0),
			CreatedAt: &timestamppb.Timestamp{
				Seconds: date.Unix(),
				Nanos:   int32(date.Nanosecond()),
			},
		}

		getUser = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Role:      "ROLE_USER",
			CreatedAt: date,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		userServiceMock userServiceMockFunc
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
				mock.GetMock.Expect(ctx, id).Return(getUser, nil)
				return mock
			},
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
				mock.GetMock.Expect(ctx, id).Return(nil, serviceError)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewUserAPI(userServiceMock)

			result, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
