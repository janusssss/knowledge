package main

import (
	"fmt"
	"io"
	"net/http"
)

// handleRequest 是处理请求的函数。
func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		io.WriteString(w, "GET request.\n")
	case http.MethodPost:
		io.WriteString(w, "Post request.\n")
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println(" 8080端口监听中........")
	http.ListenAndServe(":8080", nil)
}
