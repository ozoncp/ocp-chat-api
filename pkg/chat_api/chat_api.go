package chat_api

import (
	"context"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/pkg/errors"
)

//go:generate mockgen --source=./chat_api.go -destination=../mocks/chat_api/service_mock.go -package=chat_api

type Service interface {
	CreateChat(ctx context.Context, classroom uint64, link string) (*chat.Chat, error)
	DescribeChat(ctx context.Context, id uint64) (*chat.Chat, error)
	RemoveChat(ctx context.Context, id uint64) error
	ListChats(ctx context.Context) ([]*chat.Chat, error)
}

type ChatAPI struct {
	service Service
}

func New(service Service) *ChatAPI {
	return &ChatAPI{
		service: service,
	}
}

func (s *ChatAPI) CreateChat(ctx context.Context, req *CreateChatRequest) (*CreateChatResponse, error) {
	ch, err := s.service.CreateChat(ctx, req.ClassroomId, req.Link)
	if err != nil {
		return nil, errors.Wrap(err, "create chat")
	}

	return &CreateChatResponse{
		Id:      ch.ID,
		Message: "created successfully",
	}, nil
}

func (s *ChatAPI) DescribeChat(ctx context.Context, req *DescribeChatRequest) (*DescribeChatResponse, error) {
	ch, err := s.service.DescribeChat(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "describe chat")
	}
	return &DescribeChatResponse{
		Chat: &ChatInstance{
			Id:          ch.ID,
			ClassroomId: ch.ClassroomID,
			Link:        ch.Link,
		},
	}, nil
}

func (s *ChatAPI) RemoveChat(ctx context.Context, req *RemoveChatRequest) (*RemoveChatResponse, error) {
	err := s.service.RemoveChat(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "remove chat")
	}
	return &RemoveChatResponse{}, nil
}

func (s *ChatAPI) ListChats(ctx context.Context, req *ListChatsRequest) (*ListChatsResponse, error) {
	chats, err := s.service.ListChats(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list chats")
	}

	var chatInstances []*ChatInstance
	for _, ch := range chats {
		newChat := &ChatInstance{
			Id:          ch.ID,
			ClassroomId: ch.ClassroomID,
			Link:        ch.Link,
		}
		chatInstances = append(chatInstances, newChat)
	}

	return &ListChatsResponse{
		Packet: chatInstances,
	}, nil
}

func (s *ChatAPI) mustEmbedUnimplementedChatApiServer() {}
