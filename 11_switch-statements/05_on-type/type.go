package main

import "fmt"

//接通类型
//--通常我们打开变量的值
//--go允许您打开变量类型

type contact struct {
	greeting string
	name     string
}

//SwitchOnType适用于接口
//我们稍后将详细了解接口
func SwitchOnType(x interface{}) {
	switch x.(type) { //这是一种断言；断言，“x是这种类型”
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")

	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = contact{"Good to see you,", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}
