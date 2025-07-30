package service

import (
	"goapi/model"
	"sync"
)

func ChunkNumbers(number int) []model.Chunk {
	var chunks []model.Chunk
	start := 1
	for start <= number {
		end := start + model.ChunkSize - 1
		if end > number {
			end = number
		}
		chunks = append(chunks, model.Chunk{Start: start, End: end})
		start = end + 1
	}
	return chunks
}

func ProcessChunks(chunks []model.Chunk, ch chan int, wg *sync.WaitGroup) {
	for _, chunk := range chunks {
		wg.Add(1)
		go func(c model.Chunk) {
			defer wg.Done()
			sum := 0
			for i := c.Start; i <= c.End; i++ {
				sum += i
			}
			ch <- sum
		}(chunk)
	}
}

func SumOfChunks(ch <-chan int) int {
	total := 0
	for val := range ch {
		total += val
	}
	return total
}
