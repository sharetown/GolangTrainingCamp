package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43

	//b是int指针；
	//b指向存储int的内存地址
	//要查看该内存地址中的值，请在b前面添加一个*
	//这被称为取消引用
	//在本例中，*是运算符
}

/*输出：
43
0xc00001c098
0xc00001c098
43
*/
