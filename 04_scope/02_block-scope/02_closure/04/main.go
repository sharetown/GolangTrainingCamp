package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
闭包帮助我们限制多个函数使用的变量的范围
如果两个或多个函数能够访问同一变量，
该变量需要是包范围
*/