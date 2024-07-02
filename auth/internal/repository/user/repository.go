package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/bifidokk/awesome-chat/auth/internal/repository"
	"github.com/bifidokk/awesome-chat/auth/internal/repository/user/converter"
	modelRepository "github.com/bifidokk/awesome-chat/auth/internal/repository/user/model"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"github.com/jackc/pgx/v4/pgxpool"
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
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r repo) Create(ctx context.Context, data *desc.CreateRequest) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(data.Name, data.Email, data.Password, data.Role).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var userId int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r repo) Delete(ctx context.Context, id int64) error {
	query, args, err := sq.Delete(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)

	return err
}

func (r repo) Update(ctx context.Context, data *desc.UpdateRequest) error {
	builder := sq.Update(tableName).
		SetMap(sq.Eq{nameColumn: data.Name.Value, emailColumn: data.Email.Value}).
		Where(sq.Eq{idColumn: data.Id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)

	return err
}

func (r repo) Get(ctx context.Context, id int64) (*desc.GetResponse, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	query, args, err := builder.ToSql()

	if err != nil {
		return nil, err
	}

	var user modelRepository.User
	err = r.db.
		QueryRow(ctx, query, args...).
		Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepository(&user), nil
}
