package algorithm

import (
	"cmp"
	"reflect"
	"slices"
)

func SliceSame[C cmp.Ordered](a, b []C) bool {
	slices.Sort(a)
	slices.Sort(b)
	return reflect.DeepEqual(a, b)
}
