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

func (s *Service) CreateChat(ctx context.Context, req *CreateChatRequest) (*CreateChatResponse, error) {
	return &CreateChatResponse{
		Code:    200,
		Message: "ok norm response",
	}, nil
}

func (s *Service) DescribeChat(ctx context.Context, req *DescribeChatRequest) (*DescribeChatResponse, error) {
	return &DescribeChatResponse{
		Id:          222,
		ClassroomId: 111,
		Link:        "asdfasfsadf.com",
	}, nil
}

func (s *Service) RemoveChat(ctx context.Context, req *RemoveChatRequest) (*RemoveChatResponse, error) {
	return &RemoveChatResponse{
		Code:    200,
		Message: "delete ok",
	}, nil
}

func (s *Service) ListChats(ctx context.Context, req *ListChatsRequest) (*ListChatsResponse, error) {
	c1 := &ListChatOne{
		Id:          11,
		ClassroomId: 111,
		Link:        "asdfasfsadf.com",
	}

	c2 := &ListChatOne{
		Id:          22,
		ClassroomId: 222,
		Link:        "asdfasfsadf2.com",
	}

	c3 := &ListChatOne{
		Id:          33,
		ClassroomId: 333,
		Link:        "asdfasfsadf3.com",
	}

	chats := []*ListChatOne{c1, c2, c3}
	return &ListChatsResponse{
		Packet: chats,
	}, nil
}
