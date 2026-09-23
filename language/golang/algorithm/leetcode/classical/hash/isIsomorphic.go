package main

import "fmt"

type exampleIsIsomorphic struct {
	s      string
	t      string
	answer bool
}

func main() {
	s := []exampleIsIsomorphic{
		{"egg", "add", true},
		{"badc", "baba", false},
	}
	for _, v := range s {
		fmt.Printf("%v\n%v\n---------------\n", isIsomorphic(v.s, v.t), v.answer)
	}
}

func isIsomorphic(s string, t string) bool {
	ms, mt := make(map[byte]byte), make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		if ms[x] != 0 && ms[x] != y || mt[y] != 0 && mt[y] != x {
			return false
		}
		ms[x], mt[y] = y, x
	}
	return true
}
