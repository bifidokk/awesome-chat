package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/bifidokk/awesome-chat/auth/internal/client/db"
	"github.com/bifidokk/awesome-chat/auth/internal/model"
	"github.com/bifidokk/awesome-chat/auth/internal/repository"
	"github.com/bifidokk/awesome-chat/auth/internal/repository/user/converter"
	modelRepository "github.com/bifidokk/awesome-chat/auth/internal/repository/user/model"
)

const (
	tableName = "\"user\""

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new instance of UserRepository.
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r repo) Create(ctx context.Context, data *model.CreateUser) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(data.Name, data.Email, data.Password, data.Role).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var userID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r repo) Delete(ctx context.Context, id int64) error {
	query, args, err := sq.Delete(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)

	return err
}

func (r repo) Update(ctx context.Context, data *model.UpdateUser) error {
	builder := sq.Update(tableName).
		SetMap(sq.Eq{nameColumn: data.Name, emailColumn: data.Email}).
		Where(sq.Eq{idColumn: data.ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)

	return err
}

func (r repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	query, args, err := builder.ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepository.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)

	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepository(&user), nil
}
