package chat_api

import (
	"context"

	"github.com/pkg/errors"
)

type Service interface {
	CreateChat(ctx context.Context, id, classroom uint64, link string) error
	DescribeChat(ctx context.Context, id uint64) (string, error)
	RemoveChat(ctx context.Context, id uint64) error
	ListChats(ctx context.Context) ([]string, error)
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
	if err := s.service.CreateChat(ctx, req.Id, req.ClassroomId, req.Link); err != nil {
		return nil, errors.Wrap(err, "create chat")
	}

	return &CreateChatResponse{
		Code:    200,
		Message: "ok",
	}, nil
}

func (s *ChatAPI) DescribeChat(ctx context.Context, req *DescribeChatRequest) (*DescribeChatResponse, error) {
	return &DescribeChatResponse{
		Id:          222,
		ClassroomId: 111,
		Link:        "asdfasfsadf.com",
	}, nil
}

func (s *ChatAPI) RemoveChat(ctx context.Context, req *RemoveChatRequest) (*RemoveChatResponse, error) {
	return &RemoveChatResponse{
		Code:    200,
		Message: "delete ok",
	}, nil
}

func (s *ChatAPI) ListChats(ctx context.Context, req *ListChatsRequest) (*ListChatsResponse, error) {
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

func (s *ChatAPI) mustEmbedUnimplementedChatApiServer() {}
