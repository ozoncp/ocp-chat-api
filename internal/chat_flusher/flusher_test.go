package chat_flusher_test

import (
	"context"
	"fmt"
	cf "github.com/ozoncp/ocp-chat-api/internal/mocks/chat_flusher"

	"github.com/onsi/gomega"

	"github.com/ozoncp/ocp-chat-api/internal/chat"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-chat-api/internal/chat_flusher"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl                  *gomock.Controller
		mockFlushableChatRepo *cf.MockFlushableChatRepo
		chats2Save            []*chat.Chat
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockFlushableChatRepo = cf.NewMockFlushableChatRepo(ctrl)

		chatDeps1 := &chat.Deps{
			Id:          1,
			ClassroomId: 11,
			Link:        "http://chat1.com",
		}
		c1 := chat.New(chatDeps1)

		chatDeps2 := &chat.Deps{
			Id:          2,
			ClassroomId: 22,
			Link:        "http://chat2.com",
		}
		c2 := chat.New(chatDeps2)

		chatDep3 := &chat.Deps{
			Id:          3,
			ClassroomId: 33,
			Link:        "http://chat3.com",
		}
		c3 := chat.New(chatDep3)

		chatDep4 := &chat.Deps{
			Id:          4,
			ClassroomId: 44,
			Link:        "http://chat4.com",
		}
		c4 := chat.New(chatDep4)

		chats2Save = append(chats2Save, c1, c2, c3, c4)
	})

	JustBeforeEach(func() {
		ctrl.Finish()
	})

	AfterEach(func() {
		chats2Save = []*chat.Chat{}
	})

	Context("Flusher flush", func() {
		BeforeEach(func() {})

		JustBeforeEach(func() {
		})

		It("flush 4 chats, chunksize 2", func() {
			ctx := context.Background()
			mockFlushableChatRepo.EXPECT().AddBatch(gomock.Any(), gomock.Any()).Times(2)

			flusherDeps := chat_flusher.Deps{
				ChunkSize: 2,
			}
			myFlusher := chat_flusher.NewChatFlusher(flusherDeps)
			err := myFlusher.Flush(ctx, mockFlushableChatRepo, chats2Save)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("%+v finished", myFlusher)
		})

		It("flush 4 chats, chunksize 3", func() {
			ctx := context.Background()
			mockFlushableChatRepo.EXPECT().AddBatch(gomock.Any(), gomock.Any()).Times(2)

			flusherDeps := chat_flusher.Deps{
				ChunkSize: 3,
			}
			myFlusher := chat_flusher.NewChatFlusher(flusherDeps)
			err := myFlusher.Flush(ctx, mockFlushableChatRepo, chats2Save)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("%+v finished", myFlusher)
		})

		It("flush 4 chats, chunksize 1", func() {
			ctx := context.Background()
			mockFlushableChatRepo.EXPECT().AddBatch(gomock.Any(), gomock.Any()).Times(4)

			flusherDeps := chat_flusher.Deps{
				ChunkSize: 1,
			}
			myFlusher := chat_flusher.NewChatFlusher(flusherDeps)
			err := myFlusher.Flush(ctx, mockFlushableChatRepo, chats2Save)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("%+v finished", myFlusher)
		})
	})
})
