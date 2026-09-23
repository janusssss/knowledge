package standard

import (
	"log"
	"reflect"
)

func deepEqual() {
	examples := []struct{
		x any
		y any
		answer bool
	}{
		{[]int{1}, []int{1, 0}, false},
		{[]int{1}, []int{1}, true},
	}
	
	for _, e := range examples {
		if e.answer != reflect.DeepEqual(e.x, e.y){
			log.Fatalln("example: ", e)
		}
	}

}
