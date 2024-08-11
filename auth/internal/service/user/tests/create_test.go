package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/bifidokk/awesome-chat/auth/internal/model"
	"github.com/bifidokk/awesome-chat/auth/internal/repository"
	repositoryMocks "github.com/bifidokk/awesome-chat/auth/internal/repository/mocks"
	userService "github.com/bifidokk/awesome-chat/auth/internal/service/user"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx        context.Context
		createUser *model.CreateUser
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		repositoryError = fmt.Errorf("repository error")

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 10)
		role     = "ROLE_USER"

		createUser = &model.CreateUser{
			Name:            name,
			Email:           email,
			Password:        password,
			Role:            role,
			ConfirmPassword: password,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:        ctx,
				createUser: createUser,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, createUser).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx:        ctx,
				createUser: createUser,
			},
			want: 0,
			err:  repositoryError,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, createUser).Return(0, repositoryError)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			srv := userService.NewUserService(userRepositoryMock)

			result, err := srv.Create(tt.args.ctx, tt.args.createUser)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
