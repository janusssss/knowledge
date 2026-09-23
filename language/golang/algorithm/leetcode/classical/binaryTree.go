package classical

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树 root ，返回其最大深度。
二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：3

示例 2：
输入：root = [1,null,2]
输出：2

提示：
树中节点的数量在 [0, 104] 区间内。
-100 <= Node.val <= 100
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}

	var leftMaxDepth, rightMaxDepth int
	leftMaxDepth = maxDepth(root.Left)
	rightMaxDepth = maxDepth(root.Right)
	depth := max(leftMaxDepth, rightMaxDepth)

	return 1 + depth
}

/*
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1：
输入：p = [1,2,3], q = [1,2,3]
输出：true

示例 2：
输入：p = [1,2], q = [1,null,2]
输出：false

示例 3：
输入：p = [1,2,1], q = [1,1,2]
输出：false

提示：
两棵树上的节点数目都在范围 [0, 100] 内
-104 <= Node.val <= 104
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil || q == nil) || (p.Val != q.Val) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/*
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

示例 2：
输入：root = [2,1,3]
输出：[2,3,1]

示例 3：
输入：root = []
输出：[]

提示：
树中节点数目范围在 [0, 100] 内
-100 <= Node.val <= 100
*/
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false

提示：
树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return subIsSymmetric(root.Left, root.Right)
}

func subIsSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return subIsSymmetric(left.Left, right.Right) &&
		subIsSymmetric(left.Right, right.Left)
}

/*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1:
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

示例 2:
输入: preorder = [-1], inorder = [-1]
输出: [-1]

提示:
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder 和 inorder 均 无重复 元素
inorder 均出现在 preorder
preorder 保证 为二叉树的前序遍历序列
inorder 保证 为二叉树的中序遍历序列
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var i int
	for inorder[i] != preorder[0] {
		i++
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// 中序遍历的顺序是每次遍历左孩子，再遍历根节点，最后遍历右孩子。
// 后序遍历的顺序是每次遍历左孩子，再遍历右孩子，最后遍历根节点。
// 前序遍历的顺序是每次遍历根节点, 再遍历左子树，最后遍历根右孩子。

func buildTreeInorderPostorder(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			root.Left = buildTree(inorder[:i], postorder[:len(inorder[:i])])
			root.Right = buildTree(inorder[i+1:], postorder[len(inorder[:i]):len(postorder)-1])
			break
		}
	}

	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	nodes := []*Node{root}
	for len(nodes) > 0 {
		var temp []*Node
		for i, node := range nodes {
			if i+1 < len(nodes) {
				node.Next = nodes[i+1]
			}
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		nodes = temp
	}
	return root
}

/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var function func(root *TreeNode) *TreeNode
	function = func(node *TreeNode) *TreeNode {
		if node.Left == nil && node.Right == nil {
			return node
		}
		if node.Left == nil {
			return function(node.Right)
		}

		node.Left, node.Right = node.Right, node.Left
		temp := function(node.Right)
		if node.Left == nil {
			return temp
		}

		temp.Right, node.Left = node.Left, nil
		return function(temp.Right)
	}
	function(root)
}

/*
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
叶子节点 是指没有子节点的节点
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var function func(node *TreeNode, targetSum int, current int) int
	function = func(node *TreeNode, targetSum int, current int) int {
		current += node.Val
		if node.Left == nil && node.Right == nil {
			return current
		}

		if node.Left == nil {
			return function(node.Right, targetSum, current)
		}

		if node.Right == nil {
			return function(node.Left, targetSum, current)
		}

		result := function(node.Left, targetSum, current)
		if result == targetSum {
			return result
		}
		return function(node.Right, targetSum, current)
	}

	return targetSum == function(root, targetSum, 0)
}

/*
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：
例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。
叶节点 是指没有子节点的节点。
*/
func sumNumbers(root *TreeNode) int {
	var function func(node *TreeNode, num int) int
	function = func(node *TreeNode, num int) int {
		if node == nil {
			return 0
		}

		num = num*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return num
		}
		return function(node.Right, num) + function(node.Left, num)
	}

	return function(root, 0)
}

/*
二叉树中的路径被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。
同一个节点在一条路径序列中 至多出现一次 。该路径至少包含一个节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点root ，返回其最大路径和 。
*/
func maxPathSum(root *TreeNode) int {
	sum := math.MinInt64
	var maxGrant func(root *TreeNode) int

	maxGrant = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := max(maxGrant(root.Left), 0)
		r := max(maxGrant(root.Right), 0)
		m := root.Val + l + r

		sum = max(sum, m)
		return root.Val + max(l, r)
	}

	maxGrant(root)
	return sum
}

//BSTIterator
/*
实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。
指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
int next()将指针向右移动，然后返回指针处的数字。

注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。
*/
type BSTIterator struct {
	i *iterator
}

type iterator struct {
	current int
	values  []int
}

func Constructor(root *TreeNode) BSTIterator {
	values := []int{math.MinInt64}
	result := BSTIterator{i: &iterator{}}
	var meddle func(root *TreeNode)
	meddle = func(root *TreeNode) {
		if root == nil {
			return
		}
		meddle(root.Left)
		values = append(values, root.Val)
		meddle(root.Right)
	}
	meddle(root)

	result.i.values = values
	return result
}

func (b *BSTIterator) Next() int {
	b.i.current++
	return b.i.values[b.i.current]
}

func (b *BSTIterator) HasNext() bool {
	return b.i.current+1 < len(b.i.values)
}

/*
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
若最底层为第 h 层（从第 0 层开始），则该层包含 1~ 2h 个节点。
*/
func countNodes(root *TreeNode) int {
	var count int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		count++
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return count
}

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var first func(root, p, q *TreeNode)
	var second func(root, target *TreeNode) bool
	var result *TreeNode
	var target *TreeNode
	first = func(root, p, q *TreeNode) {
		if root == nil {
			return
		}
		if root == p || root == q {
			if root == p {
				target = q
			} else {
				target = p
			}
		}

		if target == nil {
			first(root.Left, p, q)
		}
		if target == nil {
			first(root.Right, p, q)
		}

		if result == nil && target != nil {
			if second(root, target) {
				result = root
			}
		}
	}

	second = func(root, target *TreeNode) bool {
		if root == nil {
			return false
		}
		if root == target {
			return true
		}

		return second(root.Left, target) || second(root.Right, target)
	}

	first(root, p, q)
	return result
}

/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	var deep int
	var dfs func(root *TreeNode, current int)

	dfs = func(root *TreeNode, current int) {
		if root == nil {
			return
		}

		if current > deep {
			deep++
			result = append(result, root.Val)
		}

		dfs(root.Right, current+1)
		dfs(root.Left, current+1)
	}

	dfs(root, 1)
	return result
}

/*
给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受
*/
func averageOfLevels(root *TreeNode) []float64 {
	var pool [][]int
	var collection func(root *TreeNode, depth int)
	collection = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if depth > len(pool)-1 {
			pool = append(pool, []int{})
		}
		pool[depth] = append(pool[depth], root.Val)

		collection(root.Left, depth+1)
		collection(root.Right, depth+1)
	}

	collection(root, 0)

	result := make([]float64, len(pool))
	for i, v := range pool {
		var sum int
		for _, vv := range v {
			sum += vv
		}
		result[i] = float64(sum) / float64(len(v))
	}

	return result
}

/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
*/
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(result)-1 < depth {
			result = append(result, []int{})
		}
		result[depth] = append(result[depth], root.Val)

		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}

	dfs(root, 0)
	return result
}

/*
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(result)-1 < depth {
			result = append(result, []int{})
		}

		if depth%2 == 0 {
			result[depth] = append(result[depth], root.Val)
		} else {
			result[depth] = append([]int{root.Val}, result[depth]...)
		}

		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}

	dfs(root, 0)
	return result
}

/*
给你一个二叉搜索树的根节点 root ，返回树中任意两不同节点值之间的最小差值。
差值是一个正数，其数值等于两值之差的绝对值。

二叉搜索树（Binary Search Tree, BST） 是一种特殊的二叉树，它通过节点值的排列规则实现了高效的查找、插入和删除操作。以下是它的核心特性和应用：
左子树所有节点值 < 根节点值
右子树所有节点值 > 根节点值
左右子树也必须是二叉搜索树（递归定义）
*/
func getMinimumDifference(root *TreeNode) int {
	minimum := math.MaxInt64
	pre := math.MinInt64
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if pre != math.MinInt64 {
			minimum = min(minimum, root.Val-pre)
		}
		pre = root.Val
		dfs(root.Right)
	}

	dfs(root)
	return minimum
}

/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）

提示：
树中的节点数为 n 。
1 <= k <= n <= 104
0 <= Node.val <= 104
*/
func kthSmallest(root *TreeNode, k int) int {
	var count int
	result := math.MinInt64
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		count++
		if count == k && result == math.MinInt64 {
			result = root.Val
			return
		}
		if result != math.MinInt64 {
			dfs(root.Right)
		}
	}

	dfs(root)
	return result
}

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含小于当前节点的数。
只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, maximum, minimum int) bool
	dfs = func(root *TreeNode, maximum, minimum int) bool {
		if root == nil {
			return true
		}
		if root.Val >= maximum || root.Val <= minimum {
			return false
		}
		return dfs(root.Left, root.Val, minimum) && dfs(root.Right, maximum, root.Val)
	}

	return dfs(root, math.MaxInt64, math.MinInt64)
}
