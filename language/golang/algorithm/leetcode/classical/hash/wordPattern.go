/*
给定一种规律 pattern 和一个字符串 s ，判断s是否遵循相同的规律。
这里的遵循指完全匹配，例如：pattern里的每个字母和字符串s中的每个非空单词之间存在着双向连接的对应规律。

示例1:
输入: pattern = "abba", s = "dog cat cat dog"
输出: true
示例 2:
输入:pattern = "abba", s = "dog cat cat fish"
输出: false
示例 3:
输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false

提示:
1 <= pattern.length <= 300
pattern 只包含小写英文字母
1 <= s.length <= 3000
s 只包含小写英文字母和 ' '
s 不包含任何前导或尾随对空格
s 中每个单词都被单个空格分隔
*/
package main

import (
	"go-career/algorithm/leetcode"
	"strings"
)

type exampleWordPattern struct {
	pattern string
	c       string
	answer  bool
}

func main() {
	s := []exampleWordPattern{{"abba", "dog dog dog dog", false}}
	for _, v := range s {
		leetcode.Printf(wordPattern(v.pattern, v.c), v.answer)
	}
}
func wordPattern(pattern string, s string) bool {
	sl := strings.Split(s, " ")
	if len(sl) != len(pattern) {
		return false
	}

	mx, my := make(map[byte]string), make(map[string]byte)
	for i := 0; i < len(sl); i++ {
		x, y := pattern[i], sl[i]
		if w, ok := mx[x]; ok && w != sl[i] {
			return false
		}
		if w, ok := my[y]; ok && w != pattern[i] {
			return false
		}
		mx[x], my[y] = y, x
	}
	return true
}
