package main

import "fmt"

func main() {
	fmt.Println(x) //使用前未先声明
	fmt.Println(y)
	x := 42 //声明后未使用
}

var y = 42
