package chat

import (
	"context"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	"github.com/bifidokk/awesome-chat/chat-server/internal/model"
	"github.com/bifidokk/awesome-chat/chat-server/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableName = "chat"

	idColumn        = "id"
	usernamesColumn = "usernames"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

// NewRepository creates a new instance of ChatRepository.
func NewRepository(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}

func (r repo) Create(ctx context.Context, data *model.CreateChat) (int64, error) {
	usernamesJSON, err := json.Marshal(data.Usernames)
	if err != nil {
		return 0, err
	}

	builderInsert := sq.Insert(tableName).
		Columns(usernamesColumn).
		Values(usernamesJSON).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	var chatID int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
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
