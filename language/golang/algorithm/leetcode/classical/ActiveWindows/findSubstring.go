package main

import (
	"fmt"
)

type finSubStringExample struct {
	s      string
	words  []string
	answer []int
}

func main() {
	examples := []finSubStringExample{
		{"aaaaaaaaaaaaaa", []string{"aa", "aa"}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"a", []string{"a"}, []int{0}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
	}
	for _, ex := range examples {
		fmt.Printf("calculate:\t%v\nanswer:\t\t%v\n------------------------------\n",
			findSubstring(ex.s, ex.words), ex.answer)
	}
}

/*
1.
*/
func findSubstring(s string, words []string) []int {
	size := len(words[0])
	m := len(words)
	length := size * len(words)
	n := len(s) - size
	ans := make([]int, 0)
	for i := 0; i < size; i++ {
		wordsMap := make(map[string]int)
		for _, word := range words {
			wordsMap[word]++
		}
		flag := m
		for j := i; j <= length+i-size && j <= n; j += size {
			key := s[j : j+size]
			value, ok := wordsMap[key]
			if ok {
				if value > 0 {
					flag--
				}
				wordsMap[key]--
			}
		}
		if flag == 0 {
			ans = append(ans, i)
		}
		for k := length + i; k <= n; k += size {
			start := k - length
			key := s[start : start+size]
			value, ok := wordsMap[key]
			if ok {
				if value >= 0 {
					flag++
				}
				wordsMap[key]++
			}

			key = s[k : k+size]
			value, ok = wordsMap[key]
			if ok {
				if value > 0 {
					flag--
				}
				wordsMap[key]--
			}

			if flag == 0 {
				ans = append(ans, k-length+size)
			}
		}
	}
	return ans
}
