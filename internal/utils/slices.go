package utils

import (
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	errors "github.com/pkg/errors"
	"math"
)

var ErrDuplicateVal = errors.New("value is duplicating, cannot fill correctly")

func SplitToChunks(chunkSize int, slice ...int) [][]int {
	length := len(slice)
	chunkNum := length/chunkSize + 1

	chunks := make([][]int, 0, length)
	for i := 0; i < chunkNum; i++ {
		end := int(math.Min(float64((i+1)*chunkSize), float64(length)))
		newSlice := slice[i*chunkSize : end]
		if len(newSlice) > 0 {
			chunks = append(chunks, newSlice)
		}
	}
	return chunks
}

func InverseMap(m map[int]string) (map[string]int, error) {
	inversed := make(map[string]int, len(m))
	for key, value := range m {
		if previousKey, found := inversed[value]; found {
			return nil, errors.Wrapf(ErrDuplicateVal,
				"try set map[%v] == %v but here is map[%v] == %v",
				value, key, value, previousKey)
		}
		inversed[value] = key
	}
	return inversed, nil
}

func ExcludeMembersOfList(list []int, removeUs []int) []int {
	removedMap := make(map[int]struct{})
	for _, elem := range removeUs {
		removedMap[elem] = struct{}{}
	}

	resultList := []int{}
	for _, x := range list {
		if _, filterMe := removedMap[x]; !filterMe {
			resultList = append(resultList, x)
		}
	}

	return resultList
}

func SplitChatListToChunks(chunkSize int, slice ...chat.Chat) [][]chat.Chat {
	length := len(slice)
	chunkNum := length/chunkSize + 1

	chunks := make([][]chat.Chat, 0, length)
	for i := 0; i < chunkNum; i++ {
		end := int(math.Min(float64((i+1)*chunkSize), float64(length)))
		newSlice := slice[i*chunkSize : end]
		if len(newSlice) > 0 {
			chunks = append(chunks, newSlice)
		}
	}
	return chunks
}

func ChatsMap(chats []chat.Chat) (map[uint64]chat.Chat, error) {
	chatsMap := make(map[uint64]chat.Chat, len(chats))
	for _, c := range chats {
		if _, found := chatsMap[c.ID()]; found {
			return nil, errors.Wrapf(ErrDuplicateVal, "key %v is in keys map yet", c.ID())
		}
		chatsMap[c.ID()] = c
	}
	return chatsMap, nil
}
