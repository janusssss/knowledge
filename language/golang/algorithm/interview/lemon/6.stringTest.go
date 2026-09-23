package main

import (
	"fmt"
)

var source = []string{
	"APTZvA", "BddOIt", "ctuuYn", "BCd5js", "cVCuqR", "AQynrL", "AoZ62r", "BV9DXI", "cqkYj7", "ALSKpF", "CEkB4M", "By6jE3", "Aclr2o", "cLiix5", "AClM5o", "BN36oa", "BYj4K0", "cKtPyI", "BGOn7c", "BQreVu", "B7kQ15", "BHhAY0", "cbQBTI", "A2KDsf", "AwmbeJ", "BsNdy0", "BoIVCB", "C3pHMS", "CP9Wc6", "C6vyPb", "A6BTpf", "AguFNY", "AoeaF8", "AyQ3dP", "CzlhVY", "BkFrls", "C4WncK", "ASTebw", "CTpdJi", "BtGzKA", "cWtmeT", "BgLz5G", "A9Ohfh", "ASv3qg", "A4du4s", "BstIGr", "BSIkmq", "CKxdNR", "BgCF6g", "CWkjqZ",
}

func main() {
	first := cutFirst(source)
	fmt.Println(first)
	fourth := changFourth(source)
	fmt.Println(fourth)
	m := total(source)
	fmt.Println(m)
}

// 请将source中的每个字符串去除首字母后放入新的切片中
func cutFirst(source []string) []string {
	result := make([]string, 0)
	for _, v := range source {
		result = append(result, string([]rune(v)[1:]))
	}
	return result
}

// 将source中的每个字符串的第4个字符改为 "A" 例： APTZvA => APTAvA, BddOIt => BddAIt
func changFourth(source []string) []string {
	for i, v := range source {
		temp := []rune(v)
		temp[4] = 'A'
		source[i] = string(temp)
	}
	return source
}

// 将字符根据首字母进行分组并统计首字母出现的次数
func total(source []string) map[string]int {
	result := make(map[string]int)
	for _, v := range source {
		f := string([]rune(v)[0])
		if _, ok := result[f]; ok {
			result[f] += 1
		} else {
			result[f] = 1
		}
	}
	return result
}
