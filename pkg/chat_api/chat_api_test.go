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
	// 1. Создали сервис, запустили, вызвали 2 Add, закрыли. Сейвер вызывался, данные там, postgres вызывался, данные там.
	// 2. Создали сервис, запустили, завершили. Сейвер вызывался, данных нет, в постгре тоже нет.
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

		//chatDeps1 := &chat.Deps{
		//	Id:          1,
		//	ClassroomId: 11,
		//	Link:        "http://chat1.com",
		//}
		//c1 := chat.New(chatDeps1)
		//
		//chatDeps2 := &chat.Deps{
		//	Id:          2,
		//	ClassroomId: 22,
		//	Link:        "http://chat2.com",
		//}
		//c2 := chat.New(chatDeps2)
		//
		//chatDep3 := &chat.Deps{
		//	Id:          3,
		//	ClassroomId: 33,
		//	Link:        "http://chat3.com",
		//}
		//c3 := chat.New(chatDep3)
		//
		//chatDep4 := &chat.Deps{
		//	Id:          4,
		//	ClassroomId: 44,
		//	Link:        "http://chat4.com",
		//}
		//c4 := chat.New(chatDep4)
		//
		//chats2Save = append(chats2Save, c1, c2, c3, c4)
	})

	JustBeforeEach(func() {
	})

	Context("Test Chat Service", func() {
		BeforeEach(func() {})

		JustBeforeEach(func() {
		})

		It("Add 10 objs only", func() {
			ctx := context.Background()

			for i := 0; i < 10; i++ {
				err := chatService.CreateChat(ctx, uint64(i), fmt.Sprintf("http://%dclass.com", i))
				gomega.Expect(err).To(gomega.BeNil())
			}

			//chatStorage.EXPECT().AddBatch(gomock.Any(), gomock.Any()).Times(2)]
			for i := 0; i < 10; i++ {
				chatStorage.EXPECT().Insert(gomock.Any(), uint64(i), "http://%dclass.com").Times(1)
			}
			//
			//for i := 0; i < 10; i++ {
			//	chat.EXPECT().Insert(gomock.Any(), uint64(i), "http://%dclass.com")
			//}

			fmt.Printf("%+v finished", chatService)
		})

	})

	AfterEach(func() {
		//chats2Save = []*chat.Chat{}
		ctrl.Finish()
	})

})
