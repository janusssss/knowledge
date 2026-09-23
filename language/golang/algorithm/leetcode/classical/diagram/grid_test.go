package diagram

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestNumIslands(t *testing.T) {
	examples := []struct {
		grid   [][]byte
		answer int
	}{
		{[][]byte{
			{1, 1, 1, 1, 0},
			{1, 1, 0, 1, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}, 1},
		{[][]byte{
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 1},
		}, 3},
	}
	for _, example := range examples {
		result := numIslands(example.grid)
		if result != example.answer {
			fmt.Printf("result:%d,answer:%d", result, example.answer)
		}
	}
}
func TestSolve(t *testing.T) {
	examples := []struct {
		grid   [][]byte
		answer [][]byte
	}{
		{
			grid: [][]byte{
				{'O', 'X', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'O'},
				{'X', 'O', 'X', 'O', 'X'},
				{'O', 'X', 'O', 'O', 'O'},
				{'X', 'X', 'O', 'X', 'O'}},
			answer: [][]byte{
				{'O', 'X', 'X', 'O', 'X'},
				{'X', 'X', 'X', 'X', 'O'},
				{'X', 'X', 'X', 'O', 'X'},
				{'O', 'X', 'O', 'O', 'O'},
				{'X', 'X', 'O', 'X', 'O'}}},
	}
	for _, example := range examples {
		solve(example.grid)
		for _, row := range example.grid {
			fmt.Println(string(row))
		}
		for i, _ := range example.grid {
			for j, _ := range example.grid[i] {
				if example.grid[i][j] != example.answer[i][j] {
					fmt.Printf("%d==result:%s----answer:%s\n", i, string(example.grid[i]), string(example.answer[i]))
					break
				}
			}
		}

	}
}

func TestCalcEquation(t *testing.T) {
	examples := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		answer    []float64
	}{
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			answer:    []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			equations: [][]string{{"a", "e"}, {"b", "e"}},
			values:    []float64{4.0, 3.0},
			queries:   [][]string{{"a", "b"}, {"e", "e"}, {"x", "x"}},
			answer:    []float64{4.0 / 3.0, 1.0, -1.0},
		},
	}

	for _, s := range examples {
		result := calcEquation(s.equations, s.values, s.queries)
		if !reflect.DeepEqual(result, s.answer) {
			fmt.Println(result)
			fmt.Println(s.answer)
			fmt.Println()
		}
	}
}

func TestCanFinish(t *testing.T) {
	examples := []struct {
		numCourses    int
		prerequisites [][]int
		answer        bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}, true},
		{3, [][]int{{0, 1}, {0, 2}, {1, 2}}, true},
		{3, [][]int{{1, 0}, {0, 2}, {2, 1}}, false},
		{8, [][]int{{1, 0}, {2, 6}, {1, 7}, {6, 4}, {7, 0}, {0, 5}}, true},
		{7, [][]int{{1, 0}, {0, 3}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}}, true},
	}

	for _, e := range examples {
		result := canFinish(e.numCourses, e.prerequisites)
		if result != e.answer {
			log.Fatalln("example: ", e)
		}
	}
}

func TestFindOrder(t *testing.T) {
	examples := []struct {
		numcourses    int
		prerequisites [][]int
		answer        []int
	}{
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 1, 2, 3}},
		{2, [][]int{{0, 1}}, []int{1, 0}},
		{6, [][]int{{2, 3}, {1, 2}, {0, 1}, {0, 4}, {4, 5}, {5, 1}}, []int{3, 2, 1, 5, 4, 0}},
	}

	for _, e := range examples {
		result := findOrder(e.numcourses, e.prerequisites)
		if !reflect.DeepEqual(result, e.answer) {
			log.Fatalf("example: %v, result: %v", e, result)
		}
	}
}

func TestSnakesAndLadders(t *testing.T) {
	examples := []struct {
		board  [][]int
		answer int
	}{
		{[][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 35, -1, -1, 13, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 15, -1, -1, -1, -1}}, 4},
		{[][]int{
			{-1, -1, -1},
			{-1, 9, -1},
			{-1, 8, 9}}, 1},
		{[][]int{
			{1, 1, -1},
			{1, 1, 1},
			{-1, 1, 1}}, -1},
		{[][]int{
			{-1, -1, 19, 10, -1},
			{2, -1, -1, 6, -1},
			{-1, 17, -1, 19, -1},
			{25, -1, 20, -1, -1},
			{-1, 15, -1, -1, 15}}, 2},

		{[][]int{
			{-1, 42, 12, -1, 1, -1, -1},
			{-1, -1, 5, -1, -1, 46, 44},
			{18, 22, 6, 39, -1, -1, -1},
			{-1, -1, 40, -1, 1, -1, 37},
			{49, 38, 24, -1, 14, 29, -1},
			{-1, -1, 6, -1, -1, -1, 20},
			{-1, -1, 12, 10, -1, 5, 26}}, 2},

		{[][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{1, 1, 1, -1, 13, -1},
			{1, 1, 1, 1, 1, 8},
			{-1, 8, 8, 8, 8, 8}}, -1},
	}

	for _, e := range examples {
		result := snakesAndLadders(e.board)
		if result != e.answer {
			log.Fatalf("example: %v, result: %d", e, result)
		}
	}
}

func TestMinMutation(t *testing.T) {
	examples := []struct {
		startGene string
		endGene   string
		bank      []string
		answer    int
	}{
		{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		{"AGCAAAAA", "GACAAAAA", []string{"AGTAAAAA", "GGTAAAAA", "GATAAAAA", "GACAAAAA"}, 4},
		{"AACCGGTT", "AACCGCTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
	}

	for _, e := range examples {
		if result := minMutation(e.startGene, e.endGene, e.bank); result != e.answer {
			log.Fatalf("example: %v, result: %d", e, result)
		}
	}
}

func TestLadderLength(t *testing.T) {
	examples := []struct {
		beginWord string
		endWord   string
		wordList  []string
		answer    int
	}{
		{"hot", "dog", []string{"dot", "dog"}, 3},
		{"qa", "sq", []string{
			"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar",
			"ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma",
			"re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge",
			"th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh",
			"wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be",
			"fm","ta","tb","ni","mr","pa","he","lr","sq","ye"}, 5},
		{"lost", "miss", []string{"most","mist","miss","lost","fist","fish"}, 4},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
		{"a", "c", []string{"a", "b", "c"}, 2},
	}

	for _, e := range examples {
		if result := ladderLength(e.beginWord, e.endWord, e.wordList); result != e.answer {
			log.Fatalf("example: %v, result: %d", e, result)
		}
	}
}
