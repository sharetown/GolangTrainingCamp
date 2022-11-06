package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	/*输出：
	  42 - 101010 - 2a
	  42 - 101010 - 0x2a
	  42 - 101010 - 0X2A
	  42       101010          0X2A

	  %#：备用格式：添加前导
	  %#x：为十六进制添加前导 0x（%#x）或 0X（%#X）
	  %X：十六进制，大写字母，每字节两个字符
	  参考go中文网：https://studygolang.com/articles/2644
	*/

}
