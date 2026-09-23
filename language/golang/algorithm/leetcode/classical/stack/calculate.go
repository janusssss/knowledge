package main

/*
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

示例 1：
输入：s = "1 + 1"
输出：2

示例 2：
输入：s = " 2-1 + 2 "
输出：3

示例 3：
输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23

提示：
1 <= s.length <= 3 * 105
s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
s 表示一个有效的表达式
'+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
'-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
输入中不存在两个连续的操作符
每个数字和运行的计算将适合于一个有符号的 32位 整数
*/
func calculate(s string) int {
	const (
		Plus         = '+'
		Minus        = '-'
		Blank        = ' '
		LeftBracket  = '('
		RightBracket = ')'
		Zero         = '0'
		Nine         = '9'
	)

	var (
		result  int
		element int
		sign    = 1
		opts    = []int{1}
	)

	for i, v := range s {
		switch v {
		case Plus:
			sign = opts[len(opts)-1]
		case Minus:
			sign = -opts[len(opts)-1]
		case Blank:
		case LeftBracket:
			opts = append(opts, sign)
		case RightBracket:
			opts = opts[:len(opts)-1]
		default:
			if v >= Zero && v <= Nine {
				element = element*10 + int(v-Zero)
			}
			if i == len(s)-1 || (s[i+1] < Zero || s[i+1] > Nine) {
				result += element * sign
				element = 0
			}
		}
	}
	return result
}
