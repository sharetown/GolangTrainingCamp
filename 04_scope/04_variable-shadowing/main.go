package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max现在是结果，而不是函数
}

// 不要这样做；阴影变量的错误编码实践