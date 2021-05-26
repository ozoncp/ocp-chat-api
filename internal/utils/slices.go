package utils

import (
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
