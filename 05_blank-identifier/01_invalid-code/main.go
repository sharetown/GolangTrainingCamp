package main

import "fmt"

func main() {
	a := "stored in a"
	b := "stored in b"
	fmt.Println("a - ", a)
	// b未被使用-代码无效
}
