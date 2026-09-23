package algorithm

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"
	"sync"
)

/*
给定一个字符串s，验证它是否是回文串。回文串是指正读和反读都一样的字符串。

说明：
* 		只考虑字母和数字字符，忽略字母的大小写。
* 		空字符串定义为有效的回文串。

示例 1：
输入：s = "A man, a plan, a canal: Panama"
输出：true

解释：忽略标点和空格后为 "amanaplanacanalpanama"，正读反读一致。
示例 2：

输入：s = "race a car"

输出：false
解释：忽略空格后为 "raceacar"，正读反读不一致。
*/

func Palindrome(s string) bool {
	s = regexp.MustCompile("[^a-z0-9]+").ReplaceAllString(strings.ToLower(s), "")

	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}

	return true
}

/*
安克 外包-万宝盛华

给定一个递增排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
*/
func findPosition(arr []int, n int) int {
	left, right := 0, len(arr)
	for left < right {
		middle := left + (right-left)/2
		if arr[middle] >= n {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return left
}

// 腾讯输入法-商城管理 外包 并发调用多个方法
func parallelQuery() {
	fns := []func() []int{fna, fnb}
	results, wg := make([][]int, len(fns)), &sync.WaitGroup{}
	wg.Add(len(fns))

	for i, fn := range fns {
		go func(n int) {
			defer wg.Done()

			results[i] = fn()
		}(i)
	}

	wg.Wait()

	result := make([]int, 0)
	for _, v := range results {
		result = append(result, v...)
	}

	fmt.Println(result)
}

func fna() []int {
	return []int{1, 2, 3}
}

func fnb() []int {
	return []int{4, 5, 6}
}

// 百度云 配置括号
func matchBracket(str string) bool {
	bracket := map[byte]byte{'{': '}', '[': ']', '(': ')'}
	var st []byte
	for i := 0; i < len(str); i++ {
		switch c := str[i]; c {
		case '{', '[', '(':
			st = append(st, bracket[c])
		case '}', ']', ')':
			if len(st) == 0 || st[len(st)-1] != c {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0

}

// 百度地图-打车业务 线程交叉打印1~10
func printBygroutine(n int) {
	threads, flag := make([]chan int, n), make(chan struct{})
	for i := range threads {
		threads[i] = make(chan int, 1)
	}
	threads[0] <- 1

	for i := range threads {
		go func(n int) {
			for {
				if v := <-threads[n]; v <= 10 {
					fmt.Printf("%d线程：%d\n", n+1, v)
					threads[v%len(threads)] <- v + 1
				} else {
					flag <- struct{}{}
				}
			}
		}(i)
	}

	<-flag
}

/*
百度文库
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/

func sortSquare(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right, pos := 0, n-1, n-1

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[pos] = leftSquare
			left++
		} else {
			result[pos] = rightSquare
			right--
		}
		pos--
	}
	return result
}

// 商汤科技 外包 单向链表翻转
type link struct {
	value int
	next  *link
}

func reverseLink(root *link) *link {
	var prev *link
	for root != nil {
		prev, root.next, root = root, prev, root.next
	}

	return prev
}

// 理想汽车 外包 字符串数据出现次数前k个值，次数相同按自然排序
func findK(strs []string, k int) []string {
	strMap := make(map[string]int, len(strs))
	for _, str := range strs {
		strMap[str]++
	}

	result := slices.SortedFunc(maps.Keys(strMap), func(x, y string) int {
		a, b := strMap[x], strMap[y]
		if a != b {
			return b - a
		}

		return strings.Compare(x, y)
	})

	return result[:k]
}

/*
快手
二分查找函数,在已排序的切片 arr 中查找目标值 target, 如果找到目标值，返回其索引；如果未找到，返回 -1
*/
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

/*
百度
寻找二叉树最大深度
*/

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func calculateDepth(root *treeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := calculateDepth(root.left)
	rightDepth := calculateDepth(root.right)
	return max(leftDepth, rightDepth) + 1
}

// 滴滴
// 限定启动n个协程计算m数组的和, 单个协程计算数组长度为len(m)/n, 不能整除则最后一个协程计算剩余的
func work(m []int, n int) int {
	wg := sync.WaitGroup{}
	wg.Add(n)
	result, unit := make([]int, n), len(m)/n

	for j := range n {
		i := j
		go func() {
			defer wg.Done()
			start := j * unit
			end := start + unit
			if j == n-1 {
				end = len(m)
			}
			for _, v := range m[start:end] {
				result[i] += v
			}
		}()
	}
	wg.Wait()

	sum := 0
	for _, v := range result {
		sum += v
	}
	return sum
}

// 百度
// 二叉树中序遍历，使用迭代方式
func inorder(root *treeNode) (res []int) {
	var st []*treeNode
	cur := root
	for cur != nil || len(st) > 0 {
		for cur != nil {
			st = append(st, cur)
			cur = cur.left
		}

		cur = st[len(st)-1]
		st = st[:len(st)-1]
		res = append(res, cur.val)
		cur = cur.right
	}
	return
}
