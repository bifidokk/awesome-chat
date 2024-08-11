package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/bifidokk/awesome-chat/auth/internal/repository"
	repositoryMocks "github.com/bifidokk/awesome-chat/auth/internal/repository/mocks"
	userService "github.com/bifidokk/awesome-chat/auth/internal/service/user"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		repositoryError = fmt.Errorf("repository error")

		id = gofakeit.Int64()
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
				ctx: ctx,
				id:  id,
			},
			err: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			err: repositoryError,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(repositoryError)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			srv := userService.NewUserService(userRepositoryMock)

			err := srv.Delete(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
		})
	}
}
