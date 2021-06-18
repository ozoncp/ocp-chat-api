package chat_repo

import (
	"context"
	"database/sql"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var ErrNoRowsToDelete = errors.New("no chat with this params")

type PostgresRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		DB: db,
	}
}

func (p *PostgresRepo) GetAll(ctx context.Context) ([]*chat.Chat, error) {
	logger := zerolog.Ctx(ctx).With().Logger()
	logger = logger.With().Str("component", "iteration_manager").Logger()
	logger.Info().Msg("run")

	query := `select id, classroom_id, link from index_iterations;`

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
func (p *PostgresRepo) Insert(ctx context.Context, classroomID uint64, link string) (*chat.Chat, error) {
	logger := zerolog.Ctx(ctx).With().Uint64("classroom_id", classroomID).Str("link", link).Logger()
	// fixme no actions on duplicate, need uniqueness (constraints)
	query := `INSERT INTO chats (classroom_id, link)
		VALUES ($1, $2);
	`
	rows, err := p.DB.ExecContext(ctx, query, classroomID, link)
	if err != nil {
		return nil, errors.Wrap(err, "insert")
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "get id of inserted row")
	}

	logger.Info().Msgf("successfully added. last_index = %d", id)
	ch := &chat.Chat{
		ID:          uint64(id),
		ClassroomID: classroomID,
		Link:        link,
	}
	return ch, nil
}

func (p *PostgresRepo) Describe(ctx context.Context, chatID uint64) (*chat.Chat, error) {
	logger := zerolog.Ctx(ctx).With().Uint64("chat_id", chatID).Logger()
	logger.Debug().Msg("describe")
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
