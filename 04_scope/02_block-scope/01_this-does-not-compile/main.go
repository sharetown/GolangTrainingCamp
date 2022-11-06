package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {

	//无法访问x
	//这无法编译
	fmt.Println(x)
}

/*
x属于main函数中的局部变量，其作用域只在main函数中有效，在foo中无法访问
*/
