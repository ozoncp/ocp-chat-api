package utils

import "math"

func SplitToChunks(chunkSize int, slice ...int) [][]int {
	length := len(slice)
	chunkNum := length/chunkSize + 1

	chunks := make([][]int, 0, length)
	for i := 0; i < chunkNum; i++ {
		end := int(math.Min(float64((i+1)*chunkSize), float64(length)))
		chunks = append(chunks, slice[i*chunkSize:end])
	}
	return chunks
}
