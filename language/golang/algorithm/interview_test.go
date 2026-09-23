package algorithm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestParallelQuery(t *testing.T) {
	parallelQuery()
}

func TestSortSquare(t *testing.T) {
	examples := []struct {
		nums   []int
		answer []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, e := range examples {
		result := sortSquare(e.nums)
		if !reflect.DeepEqual(result, e.answer) {
			fmt.Printf("result:%v, answer:%v\n", result, e.answer)
			os.Exit(1)
		}
	}
}

func TestReverseLink(t *testing.T) {
	examples := []struct {
		root   *link
		answer *link
	}{
		{&link{1, &link{2, &link{3, nil}}}, &link{3, &link{2, &link{1, nil}}}},
	}

	fn := func(l *link) []int {
		var result []int
		for l != nil {
			result = append(result, l.value)
			l = l.next
		}

		return result
	}

	for _, e := range examples {
		result, answer := fn(reverseLink(e.root)), fn(e.answer)
		if !reflect.DeepEqual(result, answer) {
			fmt.Printf("result:%v, answer:%v", result, answer)
			os.Exit(1)
		}
	}
}

func TestFindK(t *testing.T) {
	exsamples := []struct {
		strs   []string
		k      int
		answer []string
	}{
		{[]string{"a", "a", "b"}, 1, []string{"a"}},
		{[]string{"a", "a", "b", "c", "c"}, 2, []string{"a", "c"}},
	}

	for _, e := range exsamples {
		result := findK(e.strs, e.k)
		if !reflect.DeepEqual(result, e.answer) {
			fmt.Printf("result:%v, answer:%v", result, e.answer)
			os.Exit(1)
		}
	}

}

func TestBinarySearch(t *testing.T) {
	examples := []struct {
		arr    []int
		target int
		answer int
	}{
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 7, 3},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 12, -1},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 13, 6},
	}

	for _, e := range examples {
		result := binarySearch(e.arr, e.target)
		if result != e.answer {
			fmt.Printf("result:%v, answer: %v\n", result, e.answer)
			os.Exit(1)
		}
	}
}

func TestCalculateDepth(t *testing.T) {
	examples := []struct {
		root   *treeNode
		answer int
	}{
		{&treeNode{1, &treeNode{val: 2, left: &treeNode{val: 4}, right: &treeNode{val: 5}}, &treeNode{val: 3}}, 3},
		{&treeNode{}, 1},
	}

	for _, e := range examples {
		result := calculateDepth(e.root)
		if result != e.answer {
			fmt.Printf("result:%v, answer: %v\n", result, e.answer)
			os.Exit(1)
		}
	}
}

func TestPrintBygroutine(t *testing.T) {
	printBygroutine(5)
}

func TestMatchBracket(t *testing.T) {
	examples := []struct {
		str    string
		answer bool
	}{
		{"{}", true},
		{"[[{})", false},
		{"(({{[[", false},
		{"({)}", false},
		{"(242141)", true},
		{"{{", false},
		{"{[]}", true},
	}

	for _, e := range examples {
		result := matchBracket(e.str)
		if result != e.answer {
			fmt.Printf("result:%v, answer: %v, str:%s", result, e.answer, e.str)
			os.Exit(1)
		}
	}
}
