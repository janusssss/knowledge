package classical

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	answer := &ListNode{}
	result := answer
	for l1 != nil || l2 != nil {
		if l1 != nil {
			result.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			result.Val += l2.Val
			l2 = l2.Next
		}
		if result.Val >= 10 {
			result.Next = &ListNode{Val: 1, Next: nil}
		} else if l1 == nil && l2 == nil {
			result.Next = nil
		} else {
			result.Next = &ListNode{}
		}
		result.Val = result.Val % 10
		result = result.Next
	}
	return answer
}

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}

/*
给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。
新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。
复制链表中的指针都不应指向原链表中的节点 。
例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
返回复制链表的头节点。
用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

	val：一个表示 Node.val 的整数。
	random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。

你的代码 只 接受原链表的头节点 head 作为传入参数。

示例 1：
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

示例 2：
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]

示例 3：
输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]
*/

//func copyRandomList(head *Node) *Node {
//	var answer *Node
//	oldM := map[*Node]int{}
//	temp := head
//	newM := map[int]*Node{}
//	for i := 0; temp != nil; i++ {
//		oldM[temp] = i
//		temp = temp.Next
//		newM[i] = &Node{}
//	}
//	oldM[nil], newM[-1] = -1, nil
//	for head != nil {
//		answer = newM[oldM[head]]
//		answer.Val = head.Val
//		answer.Next = newM[oldM[head.Next]]
//		answer.Random = newM[oldM[head.Random]]
//		head = head.Next
//	}
//	return newM[0]
//}

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

示例 1：
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例 2：
输入：head = [5], left = 1, right = 1
输出：[5]
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	answer := &ListNode{}
	answer = head
	interval := make([]*ListNode, 0, right-left+1)
	var before *ListNode
	for i := 1; head != nil; i++ {
		if i == right+1 {
			break
		}
		if i == left-1 {
			before = head
		}
		if i >= left && i <= right {
			interval = append(interval, head)
		}
		head = head.Next
	}

	for i := len(interval) - 1; i >= 0; i-- {
		if left == 1 && i == len(interval)-1 {
			before = interval[i]
			answer = before
		}
		before.Next = interval[i]
		before = before.Next
	}
	if before.Next != nil {
		before.Next = head
	}
	return answer
}

/*
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

示例 2：
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	var node, answer *ListNode
	interval := make([]*ListNode, 0, k)
	for i := 1; head != nil; i++ {
		if i == k {
			node, answer = head, head
		}
		interval = append(interval, head)
		head = head.Next
		if i%k == 0 {
			temp := interval[len(interval)-1]
			interval[0].Next = nil
			for ii := len(interval) - 2; ii >= 0; ii-- {
				temp.Next = interval[ii]
				temp = temp.Next
			}
			if i != k {
				node.Next = interval[len(interval)-1]
			}
			node = interval[0]
			interval = make([]*ListNode, 0, k)
		}
	}

	if len(interval) > 0 {
		node.Next = interval[0]
	}
	return answer
}

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

/*
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

示例 1：
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

示例 2：
输入：head = [1,1,1,2,3]
输出：[2,3]

提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := &ListNode{math.MinInt, head}
	answer := node
	previous := node
	same := node.Next
	node = node.Next.Next
	deleteNode := false
	for node != nil {
		if node.Val == same.Val {
			deleteNode = true
			node = node.Next
		} else {
			if deleteNode {
				deleteNode = false
				previous.Next = node
				same = node
				node = node.Next
			} else {
				previous = previous.Next
				same = node
				node = node.Next
			}
		}
	}
	if deleteNode {
		previous.Next = node
	}
	return answer.Next
}

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

示例 2：
输入：head = [0,1,2], k = 4
输出：[2,0,1]

提示：
链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	list := make([]*ListNode, 0, math.MaxUint8)
	for ; head != nil; head = head.Next {
		list = append(list, head)
	}

	k = k % len(list)
	if k == 0 {
		return list[0]
	}

	list[len(list)-1].Next = list[0]
	previous := len(list) - k
	future := previous - 1
	list[future].Next = nil

	return list[previous]
}

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

示例 1：
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

示例 2：
输入：head = [2,1], x = 2
输出：[1,2]

提示：
链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200
*/
func partition(head *ListNode, x int) *ListNode {
	previous, future := &ListNode{}, &ListNode{}
	answer, tail := previous, future

	for ; head != nil; head = head.Next {
		if head.Val < x {
			previous.Next = head
			previous = previous.Next
		} else {
			future.Next = head
			future = future.Next
		}
	}

	previous.Next, future.Next = tail.Next, nil

	return answer.Next
}

/*
Your LRUCache object will be instantiated and called as such:
obj := Constructor(capacity);
param_1 := obj.Get(key);
obj.Put(key,value);
*/
type LRUCache struct {
	container map[int]*linked
	capacity  int
	length    int
	head      *linked
}
type linked struct {
	key      int
	val      int
	next     *linked
	previous *linked
}

func ConstructorLinked(capacity int) LRUCache {
	l := LRUCache{capacity: capacity}
	l.container = make(map[int]*linked, capacity)
	if capacity == 0 {
		return l
	}

	head := &linked{}
	head.next, head.previous = head, head
	l.head = head
	if capacity == 1 {
		return l
	}

	for i := 0; i < capacity-1; i++ {
		node := &linked{}
		node.next, node.previous = head, head.previous
		head.previous.next, head.previous = node, node
	}

	return l
}

func (l *LRUCache) Get(key int) int {
	if c, ok := l.container[key]; ok {
		l.moveHead(c)
		return c.val
	}
	return -1
}

func (l *LRUCache) moveHead(c *linked) {
	if l.head == c {
		return
	}
	c.previous.next, c.next.previous = c.next, c.previous
	c.next, c.previous = l.head, l.head.previous
	l.head.previous.next, l.head.previous = c, c
	l.head = c
}

func (l *LRUCache) Put(key int, value int) {
	if c, exists := l.container[key]; exists {
		c.val = value
		l.moveHead(c)
	} else {
		if l.length < l.capacity {
			l.length++
			l.container[key], l.head.previous.val, l.head.previous.key = l.head.previous, value, key
			l.moveHead(l.head.previous)
		} else {
			delete(l.container, l.head.previous.key)
			l.head.previous.key, l.head.previous.val = key, value
			l.container[key] = l.head.previous
			l.moveHead(l.head.previous)
		}
	}
}
