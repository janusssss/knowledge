package main

import "strings"

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	answer := make([]string, 0, len(split))
	for _, v := range split {
		switch v {
		case "", ".":
		case "..":
			if len(answer) > 0 {
				answer = answer[:len(answer)-1]
			}
		default:
			answer = append(answer, v)
		}
	}
	return "/" + strings.Join(answer, "/")
}

func main() {
	println(simplifyPath("/../"))
}
