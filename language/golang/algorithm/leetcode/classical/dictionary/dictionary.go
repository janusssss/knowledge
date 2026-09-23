package dictionary

import (
	"maps"
	"slices"
)

/*
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母在一个单词中不允许被重复使用。

输入：board = [
["o","a","a","n"],
["e","t","a","e"],
["i","h","k","r"],
["i","f","l","v"]],
words = ["oath","pea","eat","rain"]
输出：["eat","oath"]

输入：board = [
["a","b"],
["c","d"]],
words = ["abcb"]
输出：[]
*/
func findWords(board [][]byte, words []string) []string {
	var placeholder byte = '#'
	type word struct {
		words [26]*word
		end   string
	}
	root := &word{}
	for _, w := range words {
		node := root
		for _, v := range w {
			i := v - 'a'
			if node.words[i] == nil {
				node.words[i] = &word{}
			}
			node = node.words[i]
		}
		node.end = w
	}

	next := func(current [2]int) [][2]int {
		x, y := current[0], current[1]
		queue := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
		return slices.DeleteFunc(queue, func(position [2]int) bool {
			x, y := position[0], position[1]
			return x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] == placeholder
		})
	}

	var fn func(queue [][2]int, node *word)
	exist := make(map[string]struct{}, len(words))
	fn = func(queue [][2]int, node *word) {
		for len(queue) > 0 {
			positon := queue[0]
			queue = queue[1:]
			x, y := positon[0], positon[1]
			i := board[x][y] - 'a'

			if node.words[i] == nil {
				continue
			}
			if word := node.words[i].end; word != "" {
				exist[word] = struct{}{}
			}

			board[x][y] = placeholder
			fn(next(positon), node.words[i])

			board[x][y] = i + 'a'
		}
	}
	for x := range board {
		for y := range board[x] {
			fn([][2]int{{x, y}}, root)
		}
	}

	return slices.Collect(maps.Keys(exist))
}
