package main

func insert(intervals [][]int, newInterval []int) [][]int {
	for i, interval := range intervals {
		if newInterval[0] <= interval[0] {
			intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i:]...)...)
			break
		}
	}
	if len(intervals) == 0 || intervals[len(intervals)-1][0] < newInterval[0] {
		intervals = append(intervals, newInterval)
	}

	var merged [][]int
	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}
