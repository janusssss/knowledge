package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if c, ok := m[v]; ok && i-c <= k {
			return true
		}
		m[v] = i
	}
	return false
}
