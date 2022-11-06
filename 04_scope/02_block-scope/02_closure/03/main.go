package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
闭包帮助我们限制多个函数使用的变量的范围
如果两个或多个函数能够访问同一变量，
该变量需要是包范围
匿名函数
没有名称的函数
func表达式
为变量赋值func
*/