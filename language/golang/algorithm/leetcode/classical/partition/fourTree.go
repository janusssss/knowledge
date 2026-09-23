package partition

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

/*
给你一个 n * n 矩阵 grid ，矩阵由若干 0 和 1 组成。请你用四叉树表示该矩阵 grid 。
你需要返回能表示矩阵 grid 的 四叉树 的根结点。
四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：
val：储存叶子结点所代表的区域的值。1 对应 True，0 对应 False。注意，当 isLeaf 为 False 时，你可以把 True 或者 False 赋值给节点，两种值都会被判题机制 接受 。
isLeaf: 当这个节点是一个叶子结点时为 True，如果它有 4 个子节点则为 False 。

class Node {
    public boolean val;
    public boolean isLeaf;
    public Node topLeft;
    public Node topRight;
    public Node bottomLeft;
    public Node bottomRight;
}

我们可以按以下步骤为二维区域构建四叉树：
如果当前网格的值相同（即，全为 0 或者全为 1），将 isLeaf 设为 True ，将 val 设为网格相应的值，并将四个子节点都设为 Null 然后停止。
如果当前网格的值不同，将 isLeaf 设为 False， 将 val 设为任意值，然后如下图所示，将当前网格划分为四个子网格。
使用适当的子网格递归每个子节点。
*/

func construct(grid [][]int) *Node {
	var subConstruct func(grid [][]int, row1, row2, col1, col2 int) *Node
	subConstruct = func(grid [][]int, row1, row2, col1, col2 int) (node *Node) {
		//fmt.Println("row1, row2, col1, col2=", row1, row2, col1, col2)

		node = &Node{}
		if len(grid) == 0 {
			return
		}

		val := grid[row1][col1]
		for row := row1; row < row2; row++ {
			for col := col1; col < col2; col++ {
				if grid[row][col] != val {
					node.TopLeft = subConstruct(grid, row1, (row1+row2)/2, col1, (col1+col2)/2)
					node.TopRight = subConstruct(grid, row1, (row1+row2)/2, (col1+col2)/2, col2)
					node.BottomLeft = subConstruct(grid, (row1+row2)/2, row2, col1, (col1+col2)/2)
					node.BottomRight = subConstruct(grid, (row1+row2)/2, row2, (col1+col2)/2, col2)
					return
				}
			}
		}

		if val == 0 {
			node.Val = true
		}
		node.IsLeaf = true
		return
	}

	n := len(grid)
	return subConstruct(grid, 0, n, 0, n)
}
