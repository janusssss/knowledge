package main

import "slices"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		s := string(b)
		m[s] = append(m[s], str)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
