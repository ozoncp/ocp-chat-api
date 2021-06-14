package chat_service

import (
	"context"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
)

type Repo interface {
	GetAll(ctx context.Context) ([]*chat.Chat, error)
	Insert(ctx context.Context, classroomID uint64, link string) (*chat.Chat, error)
	Describe(ctx context.Context, chatID uint64) (*chat.Chat, error)
	Remove(ctx context.Context, chatID uint64) error
}

type Saver interface {
	Save(ctx context.Context, ch *chat.Chat) error
}

type Deps struct {
	StatisticsSaver Saver
	StorageRepo     Repo
	QueueRepo       Repo
}

type Service struct {
	statisticsSaver Saver
	storageRepo     Repo
	queueRepo       Repo
}

func New(deps *Deps) *Service {
	return &Service{
		statisticsSaver: deps.StatisticsSaver,
		storageRepo:     deps.StorageRepo,
		queueRepo:       deps.QueueRepo,
	}
}

func (s *Service) CreateChat(ctx context.Context, classroom uint64, link string) error {
	return nil
}

func (s *Service) DescribeChat(ctx context.Context, id uint64) (string, error) {
	return "", nil
}

func (s *Service) RemoveChat(ctx context.Context, id uint64) error {
	return nil
}

func (s *Service) ListChats(ctx context.Context) ([]string, error) {
	return []string{}, nil
}
