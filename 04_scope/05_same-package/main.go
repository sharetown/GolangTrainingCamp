package main

import "fmt"

func main() {
	fmt.Println(x)
}

/*
如果要执行该文件,直接在编译器中执行该main文件会报错
需要到05_same_package目录下执行go build .\05_same_package
此时会得到一个名为05_same_package.exe的只执行文件
直接执行该可执行文件才可以正常运行
*/