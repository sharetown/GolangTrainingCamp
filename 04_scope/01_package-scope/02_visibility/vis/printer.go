package vis

import "fmt"

// PrintVar会被导出，因为它以大写字母开头
func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
}
/*
MyName和yourName都是另一个程序文件name.go中声明的变量，但因为当前文件所属的包和name.go文件所属的包都是同一个vis，
所以在当前文件下也能访问name.go中的资源
*/
