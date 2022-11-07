package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}

/*
注意：
某些操作系统（Windows）可能无法打印i<256的字符
如果您有此问题，可以使用以下代码：
fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))
UTF-8是Go使用的文本编码方案。
UTF-8可用于1-4字节。
一个字节是8位。
[]字节处理字节，即一次仅1字节（8位）。
[]int32允许我们存储4字节的值，即4字节*8位/字节=32位。
*/
