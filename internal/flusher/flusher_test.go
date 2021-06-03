package flusher_test

import (
	"fmt"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-chat-api/internal/flusher"
	"github.com/ozoncp/ocp-chat-api/internal/message"
	"github.com/ozoncp/ocp-chat-api/internal/mocks/message_repo"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl            *gomock.Controller
		mockMessageRepo *message_repo.MockMessageRepo
		m               []*message.Message
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockMessageRepo = message_repo.NewMockMessageRepo(ctrl)

		m = append(m, &message.Message{
			Timestamp: time.Time{},
			ID:        "00",
		})
		m = append(m, &message.Message{
			Timestamp: time.Time{},
			ID:        "11",
		})
		m = append(m, &message.Message{
			Timestamp: time.Time{},
			ID:        "22",
		})
		m = append(m, &message.Message{
			Timestamp: time.Time{},
			ID:        "33",
		})
	})

	JustBeforeEach(func() {
		ctrl.Finish()
	})

	AfterEach(func() {
		m = []*message.Message{}
	})

	Context("", func() {
		BeforeEach(func() {})

		JustBeforeEach(func() {
		})

		It("flush 4 messages, chunksize 2", func() {
			messageList := m

			flusherDeps := flusher.Deps{
				ChunkSize:         2,
				MessageRepository: mockMessageRepo,
			}

			mockMessageRepo.EXPECT().AddMessagesBatch(gomock.Any()).Times(2)

			myFlusher := flusher.NewFlusherMessagesToChat(flusherDeps)
			myFlusher.Flush(messageList)
			fmt.Printf("%+v finished", myFlusher)
		})

		It("flush 4 messages, chunksize 3", func() {
			messageList := m

			flusherDeps := flusher.Deps{
				ChunkSize:         3,
				MessageRepository: mockMessageRepo,
			}

			mockMessageRepo.EXPECT().AddMessagesBatch(gomock.Any()).Times(2)

			myFlusher := flusher.NewFlusherMessagesToChat(flusherDeps)
			myFlusher.Flush(messageList)
			fmt.Printf("%+v finished", myFlusher)
		})

		It("flush 4 messages, chunksize 1", func() {
			messageList := m

			flusherDeps := flusher.Deps{
				ChunkSize:         1,
				MessageRepository: mockMessageRepo,
			}

			mockMessageRepo.EXPECT().AddMessagesBatch(gomock.Any()).Times(4)

			myFlusher := flusher.NewFlusherMessagesToChat(flusherDeps)
			myFlusher.Flush(messageList)
			fmt.Printf("%+v finished", myFlusher)
		})
	})
})
