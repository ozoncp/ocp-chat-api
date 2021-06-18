package chat_repo

import (
	"context"
	"database/sql"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
)

type PostgresRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		DB: db,
	}
}

func (p *PostgresRepo) GetAll(ctx context.Context) ([]*chat.Chat, error) {
	return []*chat.Chat{}, nil
}
func (p *PostgresRepo) Insert(ctx context.Context, classroomID uint64, link string) (*chat.Chat, error) {
	return nil, nil
}

func (p *PostgresRepo) Describe(ctx context.Context, chatID uint64) (*chat.Chat, error) {
	return nil, nil
}

func (p *PostgresRepo) Remove(ctx context.Context, chatID uint64) error {
	return nil
}
