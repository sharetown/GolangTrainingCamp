package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c        // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

/*输出：
0
1
2
*/

/*
后面引用的iota可以省略
*/