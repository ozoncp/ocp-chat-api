package chat_service

import (
	"context"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/chat_flusher"
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
	Save(ctx context.Context, ch ...*chat.Chat) error
}

//go:generate mockgen --source=./service.go -destination=../mocks/chat_saver/batch_saver_mock.go -package=chat_saver
type MessageQueueConsumer interface {
	ReadChatsBatch(ctx context.Context, batchSize int) ([]*chat.Chat, error)
}

type Deps struct {
	StorageRepoSaver Saver
	StorageRepo      Repo
	QueueConsumer    MessageQueueConsumer
}

type ChatService struct {
	storageRepoSaver     Saver
	storageRepo          Repo
	storageFlushableRepo chat_flusher.FlushableChatRepo
	queueRepo            MessageQueueConsumer
}

func New(deps *Deps) *ChatService {
	return &ChatService{
		storageRepoSaver: deps.StorageRepoSaver,
		storageRepo:      deps.StorageRepo,
		queueRepo:        deps.QueueConsumer,
	}
}

func (s *ChatService) CreateMultipleChat(ctx context.Context, classroom []uint64, link []string) error {
	if len(classroom) != len(link) {
		return errors.Wrap(errors.New("different lengths of elements arrays in multiple addition"), "init multi_create_chat")
	}

	var chats []*chat.Chat
	for i := 0; i < len(classroom); i++ {
		chats = append(chats, &chat.Chat{
			ID:          0,
			ClassroomID: classroom[i],
			Link:        link[i],
		})
	}
	err := s.storageRepoSaver.Save(ctx, chats...)
	if err != nil {
		return errors.Wrap(err, "insert to repo")
	}
	// todo very very unreliable!

	return nil
}

func (s *ChatService) CreateChat(ctx context.Context, classroom uint64, link string) (*chat.Chat, error) {
	ch, err := s.storageRepo.Insert(ctx, classroom, link)
	if err != nil {
		return nil, errors.Wrap(err, "insert to repo")
	}

	if err := s.storageRepoSaver.Save(ctx, ch); err != nil {
		return nil, errors.Wrap(err, "save to statistics")
	}

	return ch, nil
}

func (s *ChatService) DescribeChat(ctx context.Context, id uint64) (*chat.Chat, error) {
	ch, err := s.storageRepo.Describe(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "describe chat")
	}
	return ch, nil
}

func (s *ChatService) RemoveChat(ctx context.Context, id uint64) error {
	if err := s.storageRepo.Remove(ctx, id); err != nil {
		return errors.Wrap(err, "remove chat")
	}
	return nil
}

func (s *ChatService) ListChats(ctx context.Context) ([]*chat.Chat, error) {
	chatsAll, err := s.storageRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list chats")
	}
	return chatsAll, nil
}
