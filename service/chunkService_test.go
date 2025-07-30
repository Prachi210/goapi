package service

import (
	"sync"
	"testing"

	"goapi/model"
)

func TestChunkNumbers(t *testing.T) {

	tests := []struct {
		number   int
		expected []model.Chunk
	}{
		{
			number:   0,
			expected: []model.Chunk{},
		},
		{
			number: 3,
			expected: []model.Chunk{
				{Start: 1, End: 3},
			},
		},
		{
			number: 200,
			expected: []model.Chunk{
				{Start: 1, End: 100},
				{Start: 101, End: 200},
			},
		},
		{
			number: 209,
			expected: []model.Chunk{
				{Start: 1, End: 100},
				{Start: 101, End: 200},
				{Start: 201, End: 209},
			},
		},
	}

	for _, tt := range tests {
		got := ChunkNumbers(tt.number)
		if len(got) != len(tt.expected) {
			t.Errorf("ChunkNumbers(%d) got %d chunks, want %d", tt.number, len(got), len(tt.expected))
			continue
		}
		for i := range got {
			if got[i] != tt.expected[i] {
				t.Errorf("ChunkNumbers(%d) chunk %d = %+v, want %+v", tt.number, i, got[i], tt.expected[i])
			}
		}
	}
}

func TestProcessChunksAndSumOfChunks(t *testing.T) {

	number := 15
	chunks := ChunkNumbers(number)
	ch := make(chan int, len(chunks))
	var wg sync.WaitGroup

	ProcessChunks(chunks, ch, &wg)
	go func() {
		wg.Wait()
		close(ch)
	}()

	total := SumOfChunks(ch)
	expected := (number * (number + 1)) / 2

	if total != expected {
		t.Errorf("SumOfChunks = %d, want %d", total, expected)
	}
}

func TestSumOfChunks_Empty(t *testing.T) {
	ch := make(chan int)
	close(ch)
	total := SumOfChunks(ch)
	if total != 0 {
		t.Errorf("SumOfChunks(empty) = %d, want 0", total)
	}
}
