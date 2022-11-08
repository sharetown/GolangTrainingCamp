package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	fmt.Println(food)//food是被上一个if语句包裹的局部变量，这里会报错

}