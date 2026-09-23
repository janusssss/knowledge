package diagram

import (
	"slices"
)

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
*/
func numIslands(grid [][]byte) int {
	var search func(x, y int)
	search = func(x, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) || grid[x][y] != '1' {
			return
		}

		grid[x][y] = '2'
		search(x-1, y)
		search(x+1, y)
		search(x, y-1)
		search(x, y+1)
	}

	var amount int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == '1' {
				search(x, y)
				amount++
			}
		}
	}

	return amount
}

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，捕获所有被围绕的区域：
连接：一个单元格与水平或垂直方向上相邻的单元格连接。
区域：连接所有 'O' 的单元格来形成一个区域。
围绕：如果您可以用 'X' 单元格连接这个区域，并且区域中没有任何单元格位于 board 边缘，则该区域被 'X' 单元格围绕。

通过原地将输入矩阵中的所有 'O' 替换为 'X' 来 捕获被围绕的区域。你不需要返回任何值。
*/
func solve(board [][]byte) {
	var fds func(i, ii int) bool
	m := make(map[[2]int]bool)

	fds = func(i, ii int) bool {
		if i < 0 || ii < 0 || i > len(board)-1 || ii > len(board[i])-1 {
			return false
		}
		if board[i][ii] == 'X' {
			return true
		}
		if m[[2]int{i, ii}] {
			return true
		}
		m[[2]int{i, ii}] = true
		return fds(i-1, ii) && fds(i+1, ii) && fds(i, ii-1) && fds(i, ii+1)
	}

	for i, row := range board {
		for ii := range row {
			if board[i][ii] != 'X' {
				if fds(i, ii) {
					for v, _ := range m {
						board[v[0]][v[1]] = 'X'
					}
				}
				clear(m)
			}
		}
	}
}

type Node struct {
	Val       int
	Neighbors []*Node
}

// 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
func cloneGraph(node *Node) *Node {
	var cg func(node *Node) *Node
	exist := make(map[*Node]*Node)

	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if v, ok := exist[node]; ok {
			return v
		}

		n := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
		exist[node] = n

		for i, neighbor := range node.Neighbors {
			n.Neighbors[i] = cg(neighbor)
		}
		return n
	}

	return cg(node)
}

/*
给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，
其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。
如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。

注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
注意：未在等式列表中出现的变量是未定义的，因此无法确定它们的答案。
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	answer := make(map[[2]string]float64)
	var add func(key [2]string, value float64)

	add = func(key [2]string, value float64) {
		if _, ok := answer[key]; ok {
			return
		}

		answer[key] = value
		answer[[2]string{key[1], key[0]}] = 1 / value
		answer[[2]string{key[0], key[0]}] = 1
		answer[[2]string{key[1], key[1]}] = 1

		for k, a := range answer {
			if k == key {
				continue
			}
			if key[0] == k[0] {
				add([2]string{k[1], key[1]}, value/a)
			}
			if key[0] == k[1] {
				add([2]string{k[0], key[1]}, value*a)
			}
			if key[1] == k[0] {
				add([2]string{key[0], k[1]}, value*a)
			}
			if key[1] == k[1] {
				add([2]string{key[0], k[0]}, value/a)
			}
		}
	}

	for i, v := range equations {
		add([2]string{v[0], v[1]}, values[i])
	}

	res := make([]float64, len(queries))
	for i, v := range queries {
		if value, ok := answer[[2]string{v[0], v[1]}]; ok {
			res[i] = value
		} else {
			res[i] = -1.0
		}
	}

	return res
}

/*
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，
其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程ai则必须先学习课程bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：
1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对互不相同
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int][]int, len(prerequisites))
	for _, p := range prerequisites {
		m[p[0]] = append(m[p[0]], p[1])
	}

	var find func(start, n int) bool
	exist := make(map[int]struct{})
	find = func(start, n int) bool {
		p, ok := m[n]
		if !ok {
			return true
		}

		if _, ok := exist[n]; ok {
			return false
		}

		for _, v := range p {
			if start > v {
				continue
			}

			exist[n] = struct{}{}
			if !find(start, v) {
				return false
			}
			delete(exist, n)
		}

		return true
	}

	for i := range numCourses {
		if ok := find(i, i); !ok {
			return false
		}
	}

	return true
}

/*
现在你总共有numCourses门课需要选，记为 0 到 numCourses - 1。
给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：[0,1]
解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。

示例 2：
输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出：[0,2,1,3]
解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。

示例 3：
输入：numCourses = 1, prerequisites = []
输出：[0]

提示：
1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
所有[ai, bi] 互不相同
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	var dfs func(n int) bool
	nodes := make(map[int][]int, len(prerequisites))
	result := make([]int, 0, len(prerequisites))
	exist := make(map[int]struct{})
	current := 0

	for _, p := range prerequisites {
		nodes[p[0]] = append(nodes[p[0]], p[1])
	}
	dfs = func(n int) bool {
		if n < current {
			return true
		}

		if _, ok := exist[n]; ok {
			return false
		}

		exist[n] = struct{}{}
		if node, ok := nodes[n]; ok {
			for _, v := range node {
				if !dfs(v) {
					return false
				}
				delete(exist, v)
			}
		}

		if !slices.Contains(result, n) {
			result = append(result, n)
		}
		return true
	}

	for i := range numCourses {
		current = i
		if !dfs(i) {
			result = []int{}
			break
		}
	}
	return result
}

/*
给你一个大小为 n x n 的整数矩阵 board ，方格按从 1 到 n2 编号，编号遵循 转行交替方式 ，
从左下角开始 （即，从 board[n - 1][0] 开始）的每一行改变方向。

你一开始位于棋盘上的方格  1。每一回合，玩家需要从当前方格 curr 开始出发，按下述要求前进：

	选定目标方格 next ，目标方格的编号在范围 [curr + 1, min(curr + 6, n2)] 。
	    该选择模拟了掷 六面体骰子 的情景，无论棋盘大小如何，玩家最多只能有 6 个目的地。
	传送玩家：如果目标方格 next 处存在蛇或梯子，那么玩家会传送到蛇或梯子的目的地。否则，玩家传送到目标方格 next 。
	当玩家到达编号 n2 的方格时，游戏结束。

如果 board[r][c] != -1 ，位于 r 行 c 列的棋盘格中可能存在 “蛇” 或 “梯子”。那个蛇或梯子的目的地将会是 board[r][c]。
编号为 1 和 n2 的方格不是任何蛇或梯子的起点。
注意，玩家在每次掷骰的前进过程中最多只能爬过蛇或梯子一次：就算目的地是另一条蛇或梯子的起点，玩家也 不能 继续移动。

举个例子，假设棋盘是 [[-1,4],[-1,3]] ，第一次移动，玩家的目标方格是 2 。
那么这个玩家将会顺着梯子到达方格 3 ，但 不能 顺着方格 3 上的梯子前往方格 4 。
（简单来说，类似飞行棋，玩家掷出骰子点数后移动对应格数，遇到单向的路径（即梯子或蛇）可以直接跳到路径的终点，但如果多个路径首尾相连，也不能连续跳多个路径）
返回达到编号为 n2 的方格所需的最少掷骰次数，如果不可能，则返回 -1。

示例 1：
输入：board = [[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]
输出：4
解释：
首先，从方格 1 [第 5 行，第 0 列] 开始。
先决定移动到方格 2 ，并必须爬过梯子移动到到方格 15 。
然后决定移动到方格 17 [第 3 行，第 4 列]，必须爬过蛇到方格 13 。
接着决定移动到方格 14 ，且必须通过梯子移动到方格 35 。
最后决定移动到方格 36 , 游戏结束。
可以证明需要至少 4 次移动才能到达最后一个方格，所以答案是 4 。

示例 2：
输入：board = [[-1,-1],[-1,3]]
输出：1

提示：
n == board.length == board[i].length
2 <= n <= 20
board[i][j] 的值是 -1 或在范围 [1, n2] 内
编号为 1 和 n2 的方格上没有蛇或梯子
*/
func snakesAndLadders(board [][]int) int {
	n := len(board)
	var fn func(num int, steps []int) int

	fn = func(num int, steps []int) int {
		println("num", num)

		var currents []int
		positions := []int{}
		for _, s := range steps {
			x := n - 1 - (s-1)/n
			y := (s - 1) % n
			if ((s-1)/n)%2 != 0 {
				y = (n - 1) - y
			}

			println("step,x,y=", s, x, y)
			v := board[x][y]
			if v == -1 {
				currents = append(currents, s)
				positions = append(positions, []int{s + 1, s + 2, s + 3, s + 4, s + 5, s + 6}...)
			} else {
				currents = append(currents, v)
				if v >= n*n {
					return num
				}
				positions = append(positions, []int{v + 1, v + 2, v + 3, v + 4, v + 5, v + 6}...)
			}
		}
		println()

		m := slices.Max(currents)
		positions = append(positions, []int{m + 1, m + 2, m + 3, m + 4, m + 5, m + 6}...)
		slices.Sort(positions)
		positions = slices.Compact(positions)
		if slices.Equal(positions, steps) || (num > 10 && slices.Equal(positions, []int{2, 3, 4, 5, 6, 7})) {
			return -1
		}

		num++
		if n > 1 && positions[len(positions)-1] >= n*n {
			return num
		}
		return fn(num, positions)
	}

	return fn(0, []int{1})
}

/*
基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。

例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。（变化后的基因必须位于基因库 bank 中）
给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。
如果无法完成此基因变化，返回 -1 。

注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。

示例 1：
输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
输出：1

示例 2：
输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
输出：2

示例 3：
输入：start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"]
输出：3

提示：
start.length == 8
end.length == 8
0 <= bank.length <= 10
bank[i].length == 8
start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成
*/
func minMutation(startGene string, endGene string, bank []string) int {
	if !slices.Contains(bank, endGene) {
		return -1
	}

	change := func(g string) (reuslt []string) {
		for i, v := range g {
			for _, s := range []string{"A", "C", "G", "T"} {
				n := g[:i] + s + g[i+1:]
				if string(v) != s && slices.Contains(bank, string(n)) {
					reuslt = append(reuslt, n)
				}
			}
		}
		return
	}

	bankMap := make(map[string][]string, len(bank))
	for _, b := range bank {
		bankMap[b] = append(bankMap[b], change(b)...)
	}

	minStep := -1
	var fn func(gene string, step int, exist []string) int
	fn = func(gene string, step int, exist []string) int {
		if gene == endGene {
			minStep = step
			return step
		}

		step++

		for _, b := range bankMap[gene] {
			if step > minStep && minStep != -1 {
				continue
			}
			if slices.Contains(exist, b) {
				continue
			}
			if b == endGene {
				minStep = step
				break
			}
			exist = append(exist, b)
			if s := fn(b, step, exist); s == -1 {
				minStep = -1
			}
			exist = slices.DeleteFunc(exist, func(s string) bool {
				return s == b
			})
		}

		return minStep
	}

	for _, g := range change(startGene) {
		fn(g, 1, []string{g})
	}

	return minStep
}

/*
字典wordList中从单词beginWord到endWord的转换序列是一个按下述规格形成的序列beginWord -> s1 -> s2 -> ... -> sk：
每一对相邻的单词只差一个字母。
对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
sk == endWord
给你两个单词 beginWord 和 endWord 和一个字典 wordList ，
返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。

示例 1：
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：5
解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。

示例 2：
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：0
解释：endWord "cog" 不在字典中，所以无法进行转换。

提示：
1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord、endWord 和 wordList[i] 由小写英文字母组成
beginWord != endWord
wordList 中的所有字符串 互不相同
*/
// func ladderLength(beginWord string, endWord string, wordList []string) int {
// 	wordId := map[string]int{}
// 	graph := [][]int{}
// 	addWord := func(word string) int {
// 		id, has := wordId[word]
// 		if !has {
// 			id = len(wordId)
// 			wordId[word] = id
// 			graph = append(graph, []int{})
// 		}
// 		return id
// 	}
// 	addEdge := func(word string) int {
// 		id1 := addWord(word)
// 		s := []byte(word)
// 		for i, b := range s {
// 			s[i] = '*'
// 			id2 := addWord(string(s))
// 			graph[id1] = append(graph[id1], id2)
// 			graph[id2] = append(graph[id2], id1)
// 			s[i] = b
// 		}
// 		return id1
// 	}

// 	for _, word := range wordList {
// 		addEdge(word)
// 	}
// 	beginId := addEdge(beginWord)

// 	fmt.Println(graph)
// 	fmt.Println(wordId)

// 	endId, has := wordId[endWord]
// 	if !has {
// 		return 0
// 	}

// 	const inf int = math.MaxInt64
// 	dist := make([]int, len(wordId))
// 	for i := range dist {
// 		dist[i] = inf
// 	}
// 	dist[beginId] = 0
// 	queue := []int{beginId}

// 	for len(queue) > 0 {
// 		v := queue[0]
// 		queue = queue[1:]
// 		if v == endId {
// 			return dist[endId]/2 + 1
// 		}
// 		for _, w := range graph[v] {
// 			if dist[w] == inf {
// 				dist[w] = dist[v] + 1
// 				queue = append(queue, w)
// 			}
// 		}
// 	}
// 	return 0
// }
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string][]string, len(wordList))
	for _, word := range append(wordList, beginWord) {
		buff := []byte(word)
		for _, w := range wordList {
			count := 0
			for i := 0; i < len(w) && count < 2; i++ {
				if w[i] != buff[i] {
					count++
				}
			}
			if count == 1 {
				wordMap[word] = append(wordMap[word], w)
			}
		}
	}

	dist := make(map[string]int)
	dist[beginWord] = 1

	queue := []string{beginWord}
	for len(queue) > 0 {
		key,queue := queue[0],queue[1:]
		if key == endWord {
			return dist[key]
		}

		for _, w := range wordMap[key] {
			if _,ok:=dist[w];!ok{
				dist[w] = dist[key] + 1
				queue = append(queue, w)
			}
		}
	}

	return 0
}
