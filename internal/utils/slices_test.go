package utils_test

import (
	"errors"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSplitToChunks(t *testing.T) {
	type TestCase struct {
		InputSlice     []int
		ChunkSize      int
		ExpectOutSlice [][]int
	}

	tests := []TestCase{
		{
			InputSlice: []int{1, 2, 3, 4, 5, 6, 7, 8},
			ChunkSize:  3,
			ExpectOutSlice: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8},
			},
		},
		{
			InputSlice: []int{1, 2, 3},
			ChunkSize:  5,
			ExpectOutSlice: [][]int{
				{1, 2, 3},
			},
		},
		{
			InputSlice:     []int{},
			ChunkSize:      534234234234,
			ExpectOutSlice: [][]int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		gotOutSlice := utils.SplitToChunks(tt.ChunkSize, tt.InputSlice...)
		assert.Equal(t, tt.ExpectOutSlice, gotOutSlice)
	}
}

func TestInverseMap(t *testing.T) {
	type TestCase struct {
		InputMap  map[int]string
		ExpectErr error
		ExpectMap map[string]int
	}

	tests := []TestCase{
		{
			InputMap: map[int]string{
				1: "one",
				2: "two",
				3: "three",
			},
			ExpectErr: nil,
			ExpectMap: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
		},
		{
			InputMap: map[int]string{
				1:  "one",
				11: "one",
				2:  "two",
				3:  "three",
			},
			ExpectErr: utils.ErrDuplicateVal,
			ExpectMap: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
		},
		{
			InputMap:  map[int]string{},
			ExpectErr: nil,
			ExpectMap: map[string]int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		gotMap, err := utils.InverseMap(tt.InputMap)
		if tt.ExpectErr != nil {
			assert.True(t, errors.Is(err, tt.ExpectErr))
			continue
		}
		require.Equal(t, err, nil)
		assert.Equal(t, tt.ExpectMap, gotMap)
	}
}

func TestExcludeMembersOfList(t *testing.T) {
	type TestCase struct {
		InputSlice     []int
		FilterSlice    []int
		ExpectOutSlice []int
	}

	tests := []TestCase{
		{
			InputSlice:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			FilterSlice:    []int{1, 2, 3},
			ExpectOutSlice: []int{4, 5, 6, 7, 8},
		},
		{
			InputSlice:     []int{},
			FilterSlice:    []int{},
			ExpectOutSlice: []int{},
		},
		{
			InputSlice:     []int{},
			FilterSlice:    []int{1, 2, 3, 4, 5, 6},
			ExpectOutSlice: []int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		gotOutSlice := utils.ExcludeMembersOfList(tt.InputSlice, tt.FilterSlice)
		assert.Equal(t, tt.ExpectOutSlice, gotOutSlice)
	}
}

func TestSplitChatListToChunks(t *testing.T) {
	type TestCase struct {
		InputSlice     []chat.Chat
		ChunkSize      int
		ExpectOutSlice [][]chat.Chat
	}

	chatDeps0 := &chat.Deps{
		Id:          0,
		ClassroomId: 1337,
		Link:        "http://welcome_to_lowload.com",
	}

	chatDeps1 := &chat.Deps{
		Id:          0,
		ClassroomId: 2288,
		Link:        "http://welcome_to_lowload.com",
	}

	chatDeps2 := &chat.Deps{
		Id:          0,
		ClassroomId: 1111,
		Link:        "http://welcome_to_lowload.com",
	}

	chat0 := chat.New(chatDeps0)
	chat1 := chat.New(chatDeps1)
	chat2 := chat.New(chatDeps2)

	tests := []TestCase{
		{
			InputSlice: []chat.Chat{*chat0, *chat1, *chat2},
			ChunkSize:  2,
			ExpectOutSlice: [][]chat.Chat{
				{*chat0, *chat1},
				{*chat2},
			},
		},
		{
			InputSlice: []chat.Chat{*chat0, *chat1, *chat2},
			ChunkSize:  1,
			ExpectOutSlice: [][]chat.Chat{
				{*chat0},
				{*chat1},
				{*chat2},
			},
		},
		{
			InputSlice:     []chat.Chat{},
			ChunkSize:      534234234234,
			ExpectOutSlice: [][]chat.Chat{},
		},
	}

	for _, tt := range tests {
		tt := tt
		gotOutSlice := utils.SplitChatListToChunks(tt.ChunkSize, tt.InputSlice...)
		assert.Equal(t, tt.ExpectOutSlice, gotOutSlice)
	}
}
