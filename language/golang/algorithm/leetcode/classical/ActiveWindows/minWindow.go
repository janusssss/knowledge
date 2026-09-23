/*
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
	如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：
 1. 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 2. 如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

示例 2：
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
*/
package main

import (
	"fmt"
)

type ninWindow struct {
	s      string
	t      string
	answer string
}

func main() {
	examples := []ninWindow{
		{"a", "b", ""},
		{"a", "a", "a"},
		{"cabwefgewcwaefgcf", "cae", "cwae"},
		{"ADOBECODEBANC", "ABC", "BANC"},
	}
	for _, ex := range examples {
		fmt.Printf("%v\n%v\n-------\n", minWindow(ex.s, ex.t), ex.answer)
	}
}

func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	m := map[rune]int{}
	for _, v := range t {
		m[v]--
	}

	var is []int
	start, end, l, do := 0, ls-1, ls, false
	for i, v := range s {
		if c, ok := m[v]; ok {
			if c < 0 {
				lt--
			}
			m[v]++
			is = append(is, i)
			if lt == 0 {
				do = true
			}

			if do {
				for a, ii := range is {
					k := rune(s[ii])
					cc, _ := m[k]
					if cc == 0 {
						is = is[a:]
						if i-ii+1 < l {
							start, end = ii, i
							l = i - ii + 1
						}
						break
					}
					m[k]--
				}
			}
		}
	}
	if do {
		return s[start : end+1]
	}
	return ""
}
