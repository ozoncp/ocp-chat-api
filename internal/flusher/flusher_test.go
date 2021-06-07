package flusher_test

import (
	"fmt"

	"github.com/onsi/gomega"

	"github.com/ozoncp/ocp-chat-api/internal/chat"

	"github.com/ozoncp/ocp-chat-api/internal/mocks/chat_repo"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-chat-api/internal/flusher"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl            *gomock.Controller
		mockMessageRepo *chat_repo.MockChatRepo
		m               []*chat.Chat
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockMessageRepo = chat_repo.NewMockChatRepo(ctrl)

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

		m = append(m, c1, c2, c3, c4)
	})

	JustBeforeEach(func() {
		ctrl.Finish()
	})

	AfterEach(func() {
		m = []*chat.Chat{}
	})

	Context("Flusher flush", func() {
		BeforeEach(func() {})

		JustBeforeEach(func() {
		})

		It("flush 4 messages, chunksize 2", func() {
			messageList := m

			flusherDeps := flusher.Deps{
				ChunkSize:      2,
				ChatRepository: mockMessageRepo,
			}

			mockMessageRepo.EXPECT().AddBatch(gomock.Any()).Times(2)

			myFlusher := flusher.NewFlusherMessagesToChat(flusherDeps)
			err := myFlusher.Flush(messageList)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("%+v finished", myFlusher)
		})

		It("flush 4 messages, chunksize 3", func() {
			messageList := m

			flusherDeps := flusher.Deps{
				ChunkSize:      3,
				ChatRepository: mockMessageRepo,
			}

			mockMessageRepo.EXPECT().AddBatch(gomock.Any()).Times(2)

			myFlusher := flusher.NewFlusherMessagesToChat(flusherDeps)
			err := myFlusher.Flush(messageList)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("%+v finished", myFlusher)
		})

		It("flush 4 messages, chunksize 1", func() {
			messageList := m

			flusherDeps := flusher.Deps{
				ChunkSize:      1,
				ChatRepository: mockMessageRepo,
			}

			mockMessageRepo.EXPECT().AddBatch(gomock.Any()).Times(4)

			myFlusher := flusher.NewFlusherMessagesToChat(flusherDeps)
			err := myFlusher.Flush(messageList)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("%+v finished", myFlusher)
		})
	})
})
