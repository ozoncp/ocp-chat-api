package chat_repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
	"github.com/pkg/errors"
)

var (
	ErrNoRowsToDelete   = errors.New("no chat with this params")
	ErrMoreThan1Created = errors.New("more than 1 entry added, it's prohibited")
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
	logger := utils.LoggerFromCtxOrCreate(ctx).With().Logger()
	logger = logger.With().Str("component", "postgres_repo").Logger()
	logger.Info().Msg("get all chats")

	query := `select id, classroom_id, link from chats;`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "db query all chats")
	}
	defer func() {
		if err := rows.Close(); err != nil {
			logger.Err(err).Msg("close stream db query response")
		}
	}()

	var chats []*chat.Chat
	for rows.Next() {
		var idFromDB uint64
		var classroom_id uint64
		var link string

		if err = rows.Scan(&idFromDB, &classroom_id, &link); err != nil {
			return nil, errors.Wrap(err, "cannot scan chats from response")
		}

		ch := &chat.Chat{
			ID:          idFromDB,
			ClassroomID: classroom_id,
			Link:        link,
		}

		chats = append(chats, ch)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "exec context")
	}

	return chats, nil
}

func (p *PostgresRepo) AddBatch(ctx context.Context, chats []*chat.Chat) error {
	logger := utils.LoggerFromCtxOrCreate(ctx).With().Logger()
	logger = logger.With().
		Str("component", "postgres_repo").
		Uint64("classroom_id", chats[0].ClassroomID).
		Str("link", chats[0].Link).Logger()
	logger.Info().Msg("insert many")

	bracketsClassAndLink := []string{}
	values := []interface{}{}
	for i, ch := range chats {
		bracketsClassAndLink = append(bracketsClassAndLink, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		values = append(values, ch.ClassroomID)
		values = append(values, ch.Link)
	}

	query := fmt.Sprintf(`INSERT INTO chats (classroom_id, link) VALUES %s;`,
		strings.Join(bracketsClassAndLink, ", "))
	logger.Info().Msgf("query: %s, len of values: %d", query, len(values))
	_, err := p.DB.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "insert multiple")
	}

	return nil
}

func (p *PostgresRepo) Insert(ctx context.Context, classroomID uint64, link string) (*chat.Chat, error) {
	logger := utils.LoggerFromCtxOrCreate(ctx).With().Logger()
	logger = logger.With().
		Str("component", "postgres_repo").
		Uint64("classroom_id", classroomID).
		Str("link", link).Logger()
	logger.Info().Msg("insert")

	query := `INSERT INTO chats (classroom_id, link)
		VALUES ($1, $2) RETURNING id;
	`
	rows, err := p.DB.QueryContext(ctx, query, classroomID, link)
	if err != nil {
		return nil, errors.Wrap(err, "insert")
	}

	var id uint64
	i := 0
	for rows.Next() {
		if i != 0 {
			return nil, errors.Wrap(ErrMoreThan1Created, "insert entry")
		}
		if err := rows.Scan(&id); err != nil {
			return nil, errors.Wrap(err, "get id of inserted row")
		}
		i++
	}

	logger.Info().Msgf("successfully added. id = %d", id)
	ch := &chat.Chat{
		ID:          uint64(id),
		ClassroomID: classroomID,
		Link:        link,
	}
	return ch, nil
}

func (p *PostgresRepo) Describe(ctx context.Context, chatID uint64) (*chat.Chat, error) {
	logger := utils.LoggerFromCtxOrCreate(ctx).With().Logger()
	logger = logger.With().
		Str("component", "postgres_repo").
		Uint64("id", chatID).Logger()
	logger.Info().Msg("desrcibe")

	query := `select id, classroom_id, link from chats WHERE id = $1;`

	rows := p.DB.QueryRowContext(ctx, query, chatID)
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "exec context")
	}

	var id uint64
	var classroomID uint64
	var link string
	err := rows.Scan(&id, &classroomID, &link)
	if err != nil {
		return nil, errors.Wrap(ErrChatNotFound, "describe chat")
	}

	ch := &chat.Chat{
		ID:          id,
		ClassroomID: classroomID,
		Link:        link,
	}

	return ch, nil
}

func (p *PostgresRepo) Remove(ctx context.Context, chatID uint64) error {
	logger := utils.LoggerFromCtxOrCreate(ctx).With().Logger()
	logger = logger.With().
		Str("component", "postgres_repo").
		Uint64("id", chatID).Logger()
	logger.Info().Msg("remove")

	query := `DELETE FROM chats WHERE id = $1;`

	res, err := p.DB.ExecContext(ctx, query, chatID)
	if err != nil {
		return errors.Wrapf(err, " delete entry: %+v", p)
	}

	nAffected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "check affected rows")
	} else if nAffected == 0 {
		return ErrNoRowsToDelete
	}

	return nil
}
