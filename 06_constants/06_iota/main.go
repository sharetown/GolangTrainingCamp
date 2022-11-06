package main

import "fmt"

const (
	_ = iota      // 0
	b = iota * 10 // 1 * 10
	c = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
}

/*输出：
10
20
*/

/*
不用的iota变量或者像跳过的地方可以使用_下划线省略
*/
