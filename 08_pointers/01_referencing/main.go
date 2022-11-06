package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)

//上面的代码使b成为指向存储int的内存地址的指针
//b是“int指针”类型
//*int--*是类型的一部分--b是*int类型
}

/*输出：
43
0xc00001c098
0xc00001c098
*/