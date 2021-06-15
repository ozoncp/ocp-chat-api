package chat_service

import (
	"context"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/pkg/errors"
)

//go:generate mockgen --source=./service.go -destination=../mocks/chat_repo/repo_mock.go -package=chat_repo

type Repo interface {
	GetAll(ctx context.Context) ([]*chat.Chat, error)
	Insert(ctx context.Context, classroomID uint64, link string) (*chat.Chat, error)
	Describe(ctx context.Context, chatID uint64) (*chat.Chat, error)
	Remove(ctx context.Context, chatID uint64) error
}

//go:generate mockgen --source=./service.go -destination=../mocks/chat_saver/batch_saver_mock.go -package=chat_saver
type Saver interface {
	Save(ctx context.Context, ch *chat.Chat) error
}

type Deps struct {
	StatisticsSaver Saver
	StorageRepo     Repo
	QueueRepo       Repo
}

type ChatService struct {
	statisticsSaver Saver
	storageRepo     Repo
	queueRepo       Repo
}

func New(deps *Deps) *ChatService {
	return &ChatService{
		statisticsSaver: deps.StatisticsSaver,
		storageRepo:     deps.StorageRepo,
		queueRepo:       deps.QueueRepo,
	}
}

func (s *ChatService) CreateChat(ctx context.Context, classroom uint64, link string) error {
	ch, err := s.storageRepo.Insert(ctx, classroom, link)
	if err != nil {
		return errors.Wrap(err, "insert to repo")
	}

	if err := s.statisticsSaver.Save(ctx, ch); err != nil {
		return errors.Wrap(err, "save to statistics")
	}

	return nil
}

func (s *ChatService) DescribeChat(ctx context.Context, id uint64) (string, error) {
	return "", nil
}

func (s *ChatService) RemoveChat(ctx context.Context, id uint64) error {
	return nil
}

func (s *ChatService) ListChats(ctx context.Context) ([]string, error) {
	return []string{}, nil
}
