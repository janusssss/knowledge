package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：
1. 左括号必须用相同类型的右括号闭合
2. 左括号必须以正确的顺序闭合
3. 每个右括号都有一个对应的相同类型的左括号

示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([])"
输出：false

输入
s ="([)]"
输出
true

提示：
1. 1 <= s.length <= 104
2. s 仅由括号 '()[]{}' 组成
*/
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	symbolMap := map[rune]rune{')': '(', '}': '{', ']': '['}
	symbols := make(map[int]rune, len(s))
	for _, symbol := range s {
		switch symbol {
		case '(', '{', '[':
			symbols[len(symbols)+1] = symbol
		default:
			if len(symbols) == 0 || symbolMap[symbol] != symbols[len(symbols)] {
				return false
			}
			delete(symbols, len(symbols))
		}
	}
	if len(symbols) > 0 {
		return false
	}
	return true
}
func main() {
	fmt.Println(isValid("()"))
}
