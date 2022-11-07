package main

import "fmt"

func main() {
	switch "Mhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
无默认故障转移
备用是可选的
--您可以通过显式声明来指定fallthrough
--在其他语言中不需要休息
*/
