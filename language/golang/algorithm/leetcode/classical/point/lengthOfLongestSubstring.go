/*
无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
*/
package main

import "fmt"

func main() {
	examples := []lengthOfLongestSubstringExample{
		{"abba", 2},
		{"dvdf", 3},
		{" ", 1},
		{"abcabcbb", 3},
	}
	for _, ex := range examples {
		fmt.Printf("calculate:\t%v\nans:\t\t%v\n----------\n",
			lengthOfLongestSubstring(ex.s), ex.ans)
	}
}

type lengthOfLongestSubstringExample struct {
	s   string
	ans int
}

/*
1. 初始子串两端 start, end 为 0, end++ 直到值相同
*/
func lengthOfLongestSubstring(s string) int {
	start, end, ans := 0, 0, 0
	subs := make(map[byte]int)
	for end < len(s) {
		index, ok := subs[s[end]]
		if ok && index >= start {
			if temp := end - start; ans < temp {
				ans = temp
			}
			start = index + 1
		}
		subs[s[end]] = end
		end++
	}
	if temp := end - start; ans < temp {
		ans = temp
	}
	return ans
}
