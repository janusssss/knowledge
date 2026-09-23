package main

import "fmt"

func main() {
	list := []string{}
	

fmt.Printf("%p",list)
	for i := 'a'; i <= 'c'; i++ {
		list = append(list, string(i))
	}
fmt.Printf("%p",list)
	
}
