package main

import "fmt"

func main() {
	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}

/*
如果我们列出10以下的所有自然数，这些自然数是3或5的倍数，我们会得到3、5、6和9。
这些倍数之和是23。
求1000以下所有3或5的倍数之和。
*/