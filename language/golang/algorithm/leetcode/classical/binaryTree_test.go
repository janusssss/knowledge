package classical

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	fmt.Println(tree)
}

func TestFlatten(t *testing.T) {
	examples := []*TreeNode{
		{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}},
		{Val: 1, Left: &TreeNode{Val: 2}},
	}
	for _, example := range examples {
		flatten(example)
		for example != nil {
			fmt.Print(example.Val)
			example = example.Right
		}
	}
}

func TestHasPathSum(t *testing.T) {
	examples := []struct {
		root   *TreeNode
		target int
	}{
		{&TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
			Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}},
			22,
		}}

	for _, example := range examples {
		fmt.Println(hasPathSum(example.root, example.target))
	}
}

func TestSumNumbers(t *testing.T) {
	examples := []struct {
		root *TreeNode
		sum  int
	}{{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, sum: 25}}

	for _, example := range examples {
		if result := sumNumbers(example.root); example.sum != result {
			t.Errorf("Sum number is incorrect")
		}
	}
}

func TestMaxPathSum(t *testing.T) {
	examples := []struct {
		root *TreeNode
		sum  int
	}{
		{root: &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, sum: 42},
		{root: &TreeNode{Val: -1, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: -6}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: -2}}}}, sum: 12},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, sum: 3},
		{root: &TreeNode{Val: -3}, sum: -3},
		{root: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}}, sum: 48},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, sum: 6},
		{root: &TreeNode{Val: -1, Left: &TreeNode{Val: -2}}, sum: -1},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: -2}, Right: &TreeNode{Val: 3}}, sum: 4},
	}
	for _, example := range examples {
		if result := maxPathSum(example.root); example.sum != result {
			t.Errorf("Max Path Sum is incorrect\nresult:%d\nanswer:%d", result, example.sum)
		}
	}
}

func TestConstructor(t *testing.T) {
	examples := []*TreeNode{{Val: 7, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}}
	for _, example := range examples {
		c := Constructor(example)
		fmt.Println(c.Next())
		fmt.Println(c.Next())
	}
}

func TestCountNodes(t *testing.T) {
	examples := []struct {
		root *TreeNode
		sum  int
	}{
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}, sum: 6},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, sum: 4},
		{root: &TreeNode{Val: 1}, sum: 1},
		{root: nil, sum: 0},
	}
	for _, example := range examples {
		if result := countNodes(example.root); example.sum != result {
			fmt.Println(example.sum)
		}
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}}
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	examples := []struct {
		root   *TreeNode
		p      *TreeNode
		q      *TreeNode
		answer *TreeNode
	}{
		{root: root, p: root.Left, q: root.Right, answer: root},
		{root: root, p: root.Left, q: root.Left.Right.Right, answer: root.Left},
		{root: root1, p: root1.Right, q: root1.Left, answer: root1},
	}

	for _, example := range examples {
		if result := lowestCommonAncestor(example.root, example.p, example.q); result != example.answer {
			fmt.Printf("result:%d, answer:%d\n", result.Val, example.answer.Val)
		}
	}
}

func TestSomething(t *testing.T) {
	fmt.Println(0 % 2)
}
