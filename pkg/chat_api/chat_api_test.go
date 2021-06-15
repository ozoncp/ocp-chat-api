package chat_api_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozoncp/ocp-chat-api/internal/chat_service"
	"github.com/ozoncp/ocp-chat-api/internal/mocks/chat_repo"
	"github.com/ozoncp/ocp-chat-api/pkg/chat_api"
)

var _ = Describe("ChatApi", func() {
	var (
		ctrl *gomock.Controller

		chatStorage     *chat_repo.MockRepo
		chatQueue       *chat_repo.MockRepo
		statisticsSaver *chat_repo.MockSaver

		chatService chat_api.Service
	)

	// тесты на Add:
	// 1. Создали сервис, запустили, вызвали 10 Add, закрыли. Сторадж вызывался на каждый из Add'ов. Сейвер вызывался на каждый из Add'ов. Очередь молчала.
	// 2. Создали сервис, запустили, завершили. Сейвер не вызывался, Сторадж тоже не вызывался, Очередь не вызывалась.
	// 3. Создали сервис, запустили, закинули туда больше данных чем он смог переварить(100 чатов, а у него буфер на 10).
	//    Записалось 10 штук, остальные потеряны. Это 10 самые новые (с 90 по 99).
	//   В постгрю записалось все что было.

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		chatStorage = chat_repo.NewMockRepo(ctrl)
		chatQueue = chat_repo.NewMockRepo(ctrl)
		statisticsSaver = chat_repo.NewMockSaver(ctrl)

		serviceDeps := &chat_service.Deps{
			StorageRepo:     chatStorage,
			QueueRepo:       chatQueue,
			StatisticsSaver: statisticsSaver,
		}

		chatService = chat_service.New(serviceDeps)

	})

	JustBeforeEach(func() {})

	Context("Test Chat Service", func() {
		BeforeEach(func() {})

		JustBeforeEach(func() {
		})

		It("Nothing happened", func() {
			_ = context.Background()

			fmt.Printf("%+v finished", chatService)
		})

		It("Add 10 objs only", func() {
			ctx := context.Background()

			for i := 0; i < 10; i++ {
				chatStorage.EXPECT().Insert(gomock.Any(), uint64(i), fmt.Sprintf("http://%dclass.com", i)).Times(1)
			}

			statisticsSaver.EXPECT().Save(gomock.Any(), gomock.Any()).Times(10)

			for i := 0; i < 10; i++ {
				err := chatService.CreateChat(ctx, uint64(i), fmt.Sprintf("http://%dclass.com", i))
				gomega.Expect(err).To(gomega.BeNil())
			}

			fmt.Printf("%+v finished", chatService)
		})

		It("Describe", func() {
			ctx := context.Background()
			idToGet := uint64(12123)
			chatStorage.EXPECT().Describe(gomock.Any(), idToGet).Times(1)

			_, err := chatService.DescribeChat(ctx, idToGet)
			gomega.Expect(err).To(gomega.BeNil())

			fmt.Printf("%+v finished", chatService)
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})

})
