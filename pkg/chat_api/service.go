package chat_api

import (
	"context"
)

type MyService interface {
	CreateChat(ctx context.Context, id, classroom int64, link string) error
	DescribeChat(ctx context.Context, id int64) (string, error)
	RemoveChat(ctx context.Context, id int64) error
	ListChats(ctx context.Context) ([]string, error)
}

type ChatAPI struct {
	service MyService
}

func New(service MyService) *ChatAPI {
	return &ChatAPI{
		service: service,
	}
}

func (s *ChatAPI) CreateChat(ctx context.Context, req *CreateChatRequest) (*CreateChatResponse, error) {
	return &CreateChatResponse{
		Code:    200,
		Message: "ok norm response",
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
