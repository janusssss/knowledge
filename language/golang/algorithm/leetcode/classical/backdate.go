package classical

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。

示例 1：

	输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCCED"
	输出：true

示例 2：

	输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "SEE"
	输出：true

示例 3：

	输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCB"
	输出：false

提示：

	m == board.length
	n = board[i].length
	1 <= m, n <= 6
	1 <= word.length <= 15
	board 和 word 仅由大小写英文字母组成
*/
func exist(board [][]byte, word string) bool {
	var (
		fn   func(position [2]int, i int)
		flag bool
	)

	fn = func(position [2]int, i int) {
		x, y := position[0], position[1]
		// fmt.Printf("x,y,i=%v,%v,%v\n", x, y, i)
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != word[i] {
			return
		}

		value := board[x][y]
		board[x][y] = '?'

		if i == len(word)-1 {
			flag = true
			return
		}

		for _, p := range [][2]int{{x, y - 1}, {x, y + 1}, {x - 1, y}, {x + 1, y}} {
			if flag {
				return
			}
			fn(p, i+1)
		}
		board[x][y] = value
	}

	for x := range board {
		for y := range board[x] {
			if flag {
				return flag
			}
			fn([2]int{x, y}, 0)
		}
	}
	return flag
}

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]

提示：
1 <= n <= 8
*/
func generateParenthesis(n int) []string {
	length := 2 * n
	values := make([]byte, length)
	for i := range values {
		if i%2 == 0 {
			values[i] = '('
		} else {
			values[i] = ')'
		}
	}

	stake := make([]byte, 0, n)
	temp := make([]byte, 0, length)
	answer := make([]string, 0, n)
	var fn func(num int)

	fn = func(num int) {
		if num > length {
			return
		}

		if len(temp) == length && len(stake) == 0 {
			answer = append(answer, string(temp))
			return
		}

		for _, v := range []byte{'(', ')'} {
			if len(stake) == 0 && v == ')' {
				continue
			}

			if v == '(' {
				stake = append(stake, v)
			} else {
				stake = stake[:len(stake)-1]
			}
			temp = append(temp, v)

			fn(num + 1)

			temp = temp[:len(temp)-1]
			if v == ')' {
				stake = append(stake, '(')
			} else {
				stake = stake[:len(stake)-1]
			}
		}
	}

	fn(0)
	return answer
}

/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

示例 1：
输入：n = 4
输出：2
解释：如上图所示，4 皇后问题存在两个不同的解法。

示例 2：
输入：n = 1
输出：1
*/
// TODO 看到不懂斜线位置对应关系
func totalNQueens(n int) int {
	columns := make([]bool, n)        // 列上是否有皇后
	diagonals1 := make([]bool, 2*n-1) // 左上到右下是否有皇后
	diagonals2 := make([]bool, 2*n-1) // 右上到左下是否有皇后
	var backtrack func(int)
	count := 0
	backtrack = func(row int) {
		if row == n {
			count++
			return
		}
		for col, hasQueen := range columns {
			d1, d2 := row+n-1-col, row+col
			if hasQueen || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[col] = true
			diagonals1[d1] = true
			diagonals2[d2] = true
			backtrack(row + 1)
			columns[col] = false
			diagonals1[d1] = false
			diagonals2[d2] = false
		}
	}
	backtrack(0)
	return count
}

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。
你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。

示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

示例 2：
输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]

示例 3：
输入: candidates = [2], target = 1
输出: []

提示：
1 <= candidates.length <= 30
2 <= candidates[i] <= 40
candidates 的所有元素 互不相同
1 <= target <= 40
*/
func combinationSum(candidates []int, target int) [][]int {
	combination, combinations, total := []int{}, [][]int{}, 0
	var fn func(curr int)

	fn = func(curr int) {
		if total > target {
			return
		}

		if total == target {
			temp := make([]int, len(combination))
			copy(temp, combination)
			combinations = append(combinations, temp)
		}

		for i := curr; i < len(candidates); i++ {
			candidate := candidates[i]
			combination = append(combination, candidate)
			total += candidate
			fn(i)
			total -= candidate
			combination = combination[:len(combination)-1]
		}
	}

	fn(0)
	return combinations
}

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]

提示：
1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/
func permute(nums []int) [][]int {
	arr := [][]int{}

	var fn func(curr int)
	fn = func(curr int) {
		if curr == len(nums)-1 {
			t := make([]int, len(nums))
			copy(t, nums)
			arr = append(arr, t)
			return
		}

		for i := curr; i < len(nums); i++ {
			nums[curr], nums[i] = nums[i], nums[curr]
			fn(curr + 1)
			nums[curr], nums[i] = nums[i], nums[curr]
		}
	}

	fn(0)
	return arr
}

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

示例 1：
输入：n = 4, k = 2
输出：
[

	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],

]

示例 2：
输入：n = 1, k = 1
输出：[[1]]

提示：
1 <= n <= 20
1 <= k <= n
*/

func combine(n int, k int) [][]int {
	c := []int{}
	var combination [][]int
	var fn func(curr int)
	fn = func(curr int) {
		if curr > n+1 {
			return
		}

		if len(c) == k {
			t := make([]int, len(c))
			copy(t, c)
			combination = append(combination, t)
			return
		}

		c = append(c, curr)
		fn(curr + 1)

		c = c[:len(c)-1]
		fn(curr + 1)
	}
	fn(1)

	return combination
}

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = "2"
输出：["a","b","c"]

提示：
1 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。
*/
func letterCombinations(digits string) []string {
	number := map[byte][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	var fn func(digits string, i int) []string
	fn = func(digits string, i int) []string {
		n := digits[i] - '0'
		if i == len(digits)-1 {
			return number[n]
		}

		strs := fn(digits, i+1)
		result := make([]string, 0, len(strs)+4)
		for _, v := range number[n] {
			for _, str := range strs {
				result = append(result, v+str)
			}
		}

		return result
	}

	return fn(digits, 0)
}
