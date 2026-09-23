package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	m := map[string]int{
		//"1":                          1,
		//"1+1":                        2,
		//"1-(     -2)":                3,
		//"-1":                         -1,
		//"-(-1)":                      1,
		//"-(2-1-(4-3))":               0,
		//"1 + 1":                      2,
		//"2147483647":                 2147483647,
		//" 2-1 + 2 ":                  3,
		"2-4-(8+2-6+(8+4-(1)+8-10))": -15,
	}
	for k, v := range m {
		if got := calculate(k); got != v {
			t.Errorf("%v got: %v, want: %v", k, got, v)
		}
	}
}
