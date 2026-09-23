package main

import "go-career/algorithm/leetcode"

type exampleLongestConsecutive struct {
	nums   []int
	answer int
}

func main() {
	s := []exampleLongestConsecutive{
		{leetcode.ToIntSlice("[100,4,200,1,3,2]"), 4},
	}
	for _, v := range s {
		leetcode.Printf(longestConsecutive(v.nums), v.answer)
	}
}
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}
	longest := 0
	for k := range m {
		long := 0
		step, forward := 1, true
	loop:
		for _, ok := m[k]; ok; _, ok = m[k] {
			long++
			delete(m, k)
			k += step
		}
		if forward {
			forward = false
			k, step = k-long-1, -1
			goto loop
		}
		if long > longest {
			longest = long
		}
	}
	return longest
}
