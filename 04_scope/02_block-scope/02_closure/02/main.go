package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
闭包帮助我们限制多个函数使用的变量的范围
如果两个或多个函数能够访问同一变量，
该变量需要是包范围
*/
