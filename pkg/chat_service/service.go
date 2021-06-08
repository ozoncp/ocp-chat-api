package chat_service

import (
	"context"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/chat_flusher"
)

//go:generate mockgen --source=./service.go -destination=../mocks/chat_repo/repo_mock.go -package=chat_repo

type Repo interface {
	GetAll() ([]*chat.Chat, error)
	RemoveByID(messageID uint64) error
	DescribeByID(messageID uint64) (string, error)
	List() (string, error)
	Add(mess *chat.Chat) error
	AddBatch(mess []*chat.Chat) error
}

//go:generate mockgen --source=./service.go -destination=../mocks/chat_flusher/flusher_mock.go -package=chat_flusher

type Flusher interface {
	Flush(repo chat_flusher.FlushableChatRepo, chats []*chat.Chat) error
}

type Deps struct {
	StorageRepo    Repo
	StorageFlusher Flusher

	QueueRepo    Repo
	QueueFlusher Flusher
}

type Service struct {
	storageRepo    Repo
	storageFlusher Flusher

	queueRepo    Repo
	queueFlusher Flusher
}

func (s *Service) mustEmbedUnimplementedChatApiServer() {
}

func New(deps *Deps) *Service {
	return &Service{
		storageRepo:    deps.StorageRepo,
		storageFlusher: deps.StorageFlusher,
		queueRepo:      deps.QueueRepo,
		queueFlusher:   deps.QueueFlusher,
	}
}

func (s *Service) CreateChat(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return &CreateResponse{
		Code:    200,
		Message: "ok norm response",
	}, nil
}
