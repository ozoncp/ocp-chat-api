package utils_test

import (
	"github.com/ozoncp/ocp-chat-api/internal/utils"
	"github.com/stretchr/testify/assert"

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
	}

	for _, tt := range tests {
		tt := tt
		gotOutSlice := utils.SplitToChunks(tt.ChunkSize, tt.InputSlice...)
		assert.Equal(t, tt.ExpectOutSlice, gotOutSlice)
	}
}
