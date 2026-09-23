package daily

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

func TestSeparateSquaresII(t *testing.T) {
	examples := []struct {
		squares [][]int
		ans     float64
	}{
		{[][]int{{0, 0, 1}, {2, 2, 1}}, 1.0},
	}

	for _, e := range examples {
		if sum := separateSquaresII(e.squares); math.Round(sum*100000)/100000 != e.ans {
			t.Errorf("expect %v, got %v", e.ans, sum)
		}
	}
}

func TestSeparateSquares(t *testing.T) {
	examples := []struct {
		squares [][]int
		ans     float64
	}{
		{[][]int{{0, 0, 1}, {2, 2, 1}}, 1.0},
		{[][]int{{2, 5, 3}, {8, 12, 4}}, 12.875},
		{[][]int{{0, 0, 2}, {1, 1, 1}}, 1.16667},
		{[][]int{{22, 27, 9}, {10, 26, 4}, {19, 21, 8}}, 28.35714},
		{[][]int{{9, 28, 4}, {21, 29, 9}, {21, 24, 8}}, 30.7381},
		{[][]int{{9, 17, 3}, {3, 16, 2}}, 17.9},
		{[][]int{{23, 29, 3}, {28, 29, 4}}, 30.78571},
	}

	for _, e := range examples {
		if sum := separateSquares(e.squares); math.Round(sum*100000)/100000 != e.ans {
			t.Errorf("expect %v, got %v", e.ans, sum)
		}
	}
}

func TestSomething(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	nums = slices.DeleteFunc(nums, func(s int) bool {
		return true
	})

	fmt.Println(nums)
}

func TestMinimumDeleteSum(t *testing.T) {
	examples := []struct {
		s1, s2 string
		ans    int
	}{
		{"sea", "eat", 231},
		{"delete", "leet", 403},
	}
	fmt.Println("s=", 's')
	fmt.Println("e", 'e')
	fmt.Println("a", 'a')
	fmt.Println("t", 't')
	/*
					e	a	t
		   		0	101	198	415
		   s	115	216	313	530
		   e	216	115	152	268
		   a	253	152	115	231
	*/
	for _, e := range examples {
		if sum := minimumDeleteSum(e.s1, e.s2); sum != e.ans {
			t.Errorf("expect %d, got %d", e.ans, sum)
		}
	}
}
