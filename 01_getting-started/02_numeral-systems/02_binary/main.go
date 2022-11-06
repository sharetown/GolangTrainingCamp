package main

import "fmt"

func main() {
	fmt.Printf("%d - %b \n", 42, 42)
	fmt.Printf("%d - %o \n", 42, 42)
	fmt.Printf("%d - %x \n", 42, 42)

	/*输出：
	    42 - 101010
	42 - 52
	42 - 2a

	fmt.Printf()可以通过占位符按格式打印数据
	%d表示十进制
	%b表示二进制
	%o表示八进制
	%x表示十六进制
	*/

}
