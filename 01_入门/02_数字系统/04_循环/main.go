package main

import "fmt"

func main() {
	//最基本的循环
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	//输出：0 1 2 3 4 5 6 7 8 9

	//变换格式后的循环
	j := 0
	for j < 10 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()
	//输出：0 1 2 3 4 5 6 7 8 9

	//死循环
	j = 0
	for {
		j++
		if j > 10 {
			break //跳出循环
		}
		fmt.Printf("%d ", j)
	}
	fmt.Println()
	//输出：1 2 3 4 5 6 7 8 9 10

	//range循环
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for k, v := range list {
		fmt.Printf("key:%d,val:%d\n", k, v)
	}
	/*输出：
	key:0,val:1
	key:1,val:2
	key:2,val:3
	key:3,val:4
	key:4,val:5
	key:5,val:6
	key:6,val:7
	key:7,val:8
	key:8,val:9
	key:9,val:10
	*/
}
