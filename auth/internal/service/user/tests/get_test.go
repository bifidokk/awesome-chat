package tests

import (
	"context"
	"database/sql"
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

func TestGet(t *testing.T) {
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

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = "ROLE_USER"
		date  = gofakeit.Date()

		getUser = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: date,
			UpdatedAt: sql.NullTime{
				Valid: true,
				Time:  date,
			},
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: getUser,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(getUser, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: nil,
			err:  repositoryError,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, repositoryError)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			srv := userService.NewUserService(userRepositoryMock)

			result, err := srv.Get(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
