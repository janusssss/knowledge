package main

type node struct {
	val   int
	lefe  *node
	right *node
}

func main() {
}

func find(root *node) (result [][]int) {
	queue := []*node{root}

	for len(queue) > 0 {
		n := len(queue)
		result = append(result, []int{})
		for i := 0; i < n; i++ {
			result[len(result)-1] = append(result[len(result)-1], queue[i].val)
			if queue[i].lefe != nil {
				queue = append(queue)
			}
			if queue[i].right != nil {
				queue = append(queue)
			}

		}

	}
}
