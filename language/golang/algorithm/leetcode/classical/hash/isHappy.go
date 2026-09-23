package main

import "go-career/algorithm/leetcode"

type exampleIsHappy struct {
	n      int
	answer bool
}

func main() {
	s := []exampleIsHappy{
		{2, false},
		{19, true},
	}
	for _, item := range s {
		leetcode.Printf(isHappy(item.n), item.answer)
	}
}
func isHappy(n int) bool {
	m := make(map[int]struct{})
	for ok := false; !ok; _, ok = m[n] {
		m[n] = struct{}{}
		var nn int
		for nn = 0; n/1 != 0; n /= 10 {
			nn += (n % 10) * (n % 10)
			n = n / 10
		}
		n = nn
		if n == 1 {
			return true
		}
	}
	return false
}
