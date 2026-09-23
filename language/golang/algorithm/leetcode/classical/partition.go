package classical

import (
	"fmt"
	"slices"
)

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[

	1->4->5,
	1->3->4,
	2->6

]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：

输入：lists = []
输出：[]

示例 3：

输入：lists = [[]]
输出：[]

提示：

	k == lists.length
	0 <= k <= 10^4
	0 <= lists[i].length <= 500
	-10^4 <= lists[i][j] <= 10^4
	lists[i] 按 升序 排列
	lists[i].length 的总和不超过 10^4
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var combine func(head, tail *ListNode) *ListNode
	combine = func(head, tail *ListNode) *ListNode {
		if head == nil && tail == nil {
			return nil
		}
		if head == nil {
			return tail
		}
		if tail == nil {
			return head
		}

		root := &ListNode{Next: head}
		pre := root
		for {
			if head == nil || head.Val > tail.Val {
				pre.Next = combine(tail, head)
				break
			}
			head, pre = head.Next, head
		}

		return root.Next
	}

	var head *ListNode
	for _, node := range lists {
		head = combine(head, node)
	}
	return head
}

func printNode(root *ListNode) {
	if root == nil {
		fmt.Println()
		return
	}
	fmt.Printf("%d ", root.Val)
	printNode(root.Next)
}

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

示例 1：

	输入：head = [4,2,1,3]
	输出：[1,2,3,4]

示例 2：

	输入：head = [-1,5,3,4,0]
	输出：[-1,0,3,4,5]

示例 3：

	输入：head = []
	输出：[]

提示：

	链表中节点的数目在范围 [0, 5 * 104] 内
	-105 <= Node.val <= 105
*/

func sortList(head *ListNode) *ListNode {
	var sort func(start, end *ListNode) *ListNode
	var merge func(current, next *ListNode) *ListNode
	merge = func(current, next *ListNode) *ListNode {
		fmt.Printf("current:%v, next:%v\n", current, next)
		maximum := next
		if next == nil {
			return current
		}
		if current == nil {
			return next
		}
		if current.Val > next.Val {
			current.Next, next.Next = next.Next, current
			maximum = current
		}
		return maximum
	}
	sort = func(start, end *ListNode) *ListNode {
		fmt.Printf("start:%v, end:%v\n", start, end)
		slow, fast := start, start
		if end == start.Next {
			slow = end
			return merge(start, end)
		}
		for fast.Next != nil {
			slow = slow.Next
			fast = fast.Next
			if fast.Next != nil {
				fast = fast.Next
			}
		}

		return merge(sort(start, slow), sort(slow, end))
	}

	sort(head, nil)
	return head
}

// 作弊取巧的方法
func sortListBySort(head *ListNode) *ListNode {
	list := make([]int, 0)
	for temp := head; temp != nil; temp = temp.Next {
		list = append(list, temp.Val)
	}

	slices.Sort(list)
	for i, temp := 0, head; i < len(list); i, temp = i+1, temp.Next {
		temp.Val = list[i]
	}

	return head
}

/*
给你一个整数数组 nums ，其中元素已经按升序排列，请你将其转换为一棵二叉搜索树。

示例 1：
	输入：nums = [-10,-3,0,5,9]
	输出：[0,-3,9,-10,null,5]
	解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

示例 2：
	输入：nums = [1,3]
	输出：[3,1]
	解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。

提示：
    1 <= nums.length <= 104
    -104 <= nums[i] <= 104
    nums 按 严格递增 顺序排列
*/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	middle := len(nums) / 2
	root := &TreeNode{Val: nums[middle]}
	root.Left = sortedArrayToBST(nums[:middle])
	root.Right = sortedArrayToBST(nums[middle+1:])

	return root
}
