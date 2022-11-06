package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)
}

/*输出；
a -  43
a's memory address -  0xc00001c098
824633835672
*/
