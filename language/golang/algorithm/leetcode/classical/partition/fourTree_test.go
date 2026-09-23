package partition

import "testing"

func TestConstruct(t *testing.T) {
	samples := []struct {
		grid [][]int
	}{
		{grid: [][]int{
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0}}},
		{grid: [][]int{{0, 1}, {1, 0}}},
	}

	for _, sample := range samples {
		construct(sample.grid)
	}
}
