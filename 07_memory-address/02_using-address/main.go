package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("输入米单位值: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " 米是 ", yards, " 码.")
}