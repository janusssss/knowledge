package daily

import (
	"slices"
)

/*
给你一个二维整数数组 squares ，其中 squares[i] = [xi, yi, li] 表示一个与 x 轴平行的正方形的左下角坐标和正方形的边长。
找到一个最小的 y 坐标，它对应一条水平线，该线需要满足它以上正方形的总面积 等于 该线以下正方形的总面积。

答案如果与实际答案的误差在 10-5 以内，将视为正确答案。
注意：正方形可能会重叠。重叠区域只统计一次。

示例 1：

	输入： squares = [[0,0,1],[2,2,1]]
	输出： 1.00000
	解释：任何在 y = 1 和 y = 2 之间的水平线都会有 1 平方单位的面积在其上方，1 平方单位的面积在其下方。最小的 y 坐标是 1。

示例 2：

	输入： squares = [[0,0,2],[1,1,1]]
	输出： 1.00000
	解释：由于蓝色正方形和红色正方形有重叠区域且重叠区域只统计一次。所以直线 y = 1 将正方形分割成两部分且面积相等。

提示：

	1 <= squares.length <= 5 * 104
	squares[i] = [xi, yi, li]
	squares[i].length == 3
	0 <= xi, yi <= 109
	1 <= li <= 109
	所有正方形的总面积不超过 1015。
*/
func separateSquaresII(squares [][]int) float64 {
	slices.SortFunc(squares, func(a, b []int) int {
		if a[1] == b[1] {
			return a[2] - b[2]
		}
		return a[1] - b[1]
	})

	calculate := func(i, curr int) (nextI int, up, area int) {
		currRow := squares[i]
		next := currRow[1] + currRow[2]
		nextI++
		for j := i + 1; j < len(squares); j++ {
			if currRow[1] == squares[j][1] {
				continue
			}

			if up = currRow[1] + currRow[2]; squares[j][1] <= up {
				next, nextI = up, j
			} else {
				next, nextI = squares[j][1], j-1
			}
			break
		}

		xs := make([][2]int, 0)
		for j := i; squares[i][1] != curr; i++ {
			flag := false
			left, right := squares[j][0], squares[j][0]+squares[j][2]
			for k := 0; k < len(xs); k++ {
				if right < xs[k][0] || left > xs[k][1] {
					continue
				}
				flag = true
				xs[k][0], xs[k][1] = min(xs[k][0], left), max(xs[k][1], right)
				k--
			}
			if flag {
				xs = append(xs, [2]int{left, right})
			}
		}

		for _, x := range xs {
			area += (x[1] - x[0]) * (next - curr)
		}

		return
	}

	var curr, area, total int
	for i := 0; i < len(squares); {
		i, area, curr = calculate(i, curr)
		total += area
	}
	half := total / 2
	total = 0

	for i := 0; i < len(squares); {
		before := curr
		i, area, curr = calculate(i, curr)
		total += area
		if total == half {
			return float64(curr)
		}
		if total > half {
			return float64(total-area-half)/float64(area)*float64(curr-before) + float64(before)
		}
	}

	return 0
}

/*
给你一个二维整数数组 squares ，其中 squares[i] = [xi, yi, li] 表示一个与 x 轴平行的正方形的左下角坐标和正方形的边长。
找到一个最小的 y 坐标，它对应一条水平线，该线需要满足它以上正方形的总面积 等于 该线以下正方形的总面积。

答案如果与实际答案的误差在 10-5 以内，将视为正确答案。
注意：正方形 可能会 重叠。重叠区域应该被 多次计数 。

示例 1：
输入： squares = [[0,0,1],[2,2,1]]
输出： 1.00000

示例 1：
输入： squares = [[0,0,2],[1,1,1]]
输出： 1.16667
*/
func separateSquares(squares [][]int) float64 {
	maxY, half := 0.0, 0.0
	for _, row := range squares {
		half += float64(row[2] * row[2])
		if y := float64(row[1] + row[2]); y > maxY {
			maxY = y
		}
	}
	half, precision := half/2, 1e-6

	bottom, top := 0.0, maxY
	for top-bottom > precision {
		limit, temp := (bottom+top)/2, 0.0
		for _, row := range squares {
			if y, h := float64(row[1]), float64(row[2]); limit > y {
				temp += h * min(limit-y, h)
			}
		}

		if temp >= half {
			top = limit
		} else {
			bottom = limit
		}
	}

	return bottom
}

/*
给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。

示例 1:
输入: s1 = "sea", s2 = "eat"
输出: 231
解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
在 "eat" 中删除 "t" 并将 116 加入总和。
结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。

示例 2:
输入: s1 = "delete", s2 = "leet"
输出: 403
解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。

提示:
0 <= s1.length, s2.length <= 1000
s1 和 s2 由小写英文字母组成
*/

func minimumDeleteSum(s1 string, s2 string) int {
	ph := make([][]int, len(s1)+1)
	for row := range ph {
		ph[row] = make([]int, len(s2)+1)
		for col := range ph[row] {
			if col == 0 && row == 0 {
				ph[row][col] = 0
			} else if row == 0 {
				ph[row][col] = int(s2[col-1]) + ph[row][col-1]
			} else if col == 0 {
				ph[row][col] = int(s1[row-1]) + ph[row-1][col]
			} else {
				if s1[row-1] == s2[col-1] {
					ph[row][col] = ph[row-1][col-1]
				} else {
					up := ph[row-1][col] + int(s1[row-1])
					left := ph[row][col-1] + int(s2[col-1])
					ph[row][col] = min(up, left)
				}
			}
		}
	}

	return ph[len(s1)][len(s2)]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个根为 root 的二叉树，每个节点的深度是该节点到根的最短距离 。
返回包含原始树中所有最深节点的最小子树 。
如果一个节点在 整个树 的任意节点之间具有最大的深度，则该节点是 最深的 。
一个节点的 子树 是该节点加上它的所有后代的集合。

提示：
树中节点的数量在 [1, 500] 范围内。
0 <= Node.val <= 500
每个节点的值都是 独一无二 的
*/
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var find func(root *TreeNode) (deep int, subtree *TreeNode)
	find = func(root *TreeNode) (deep int, subtree *TreeNode) {
		if root == nil {
			return 0, nil
		}

		leftDeep, leftSubtree := find(root.Left)
		rightDeep, rightSubtree := find(root.Right)

		if leftDeep == 0 && rightDeep == 0 {
			return 1, root
		}

		subtrees := make([]*TreeNode, 0, 3)
		deep = leftDeep
		if leftDeep == rightDeep {
			subtrees = append(subtrees, leftSubtree, rightSubtree, root)
		} else if leftDeep > rightDeep {
			subtrees = append(subtrees, leftSubtree)
		} else {
			deep = rightDeep
			subtrees = append(subtrees, rightSubtree)
		}

		slices.SortFunc(subtrees, func(a, b *TreeNode) int {
			return a.Val - b.Val
		})

		return deep + 1, subtrees[0]
	}

	_, subroot := find(root)
	return subroot
}
