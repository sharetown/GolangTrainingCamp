package main

import "fmt"

func main() {

	a := 10                    //int
	b := "golang"              //string
	c := 4.17                  //float64
	d := true                  //bool
	e := "Hello"               //string
	f := `Do you like my hat?` //string
	g := 'M'                   //int32

	fmt.Printf("%v type of %T\n", a, a)
	fmt.Printf("%v type of %T\n", b, b)
	fmt.Printf("%v type of %T\n", c, c)
	fmt.Printf("%v type of %T\n", d, d)
	fmt.Printf("%v type of %T\n", e, e)
	fmt.Printf("%v type of %T\n", f, f)
	fmt.Printf("%v type of %T\n", g, g)

	/*输出：
	10 type of int
	golang type of string
	4.17 type of float64
	true type of bool
	Hello type of string
	Do you like my hat? type of string
	77 type of int32

    1、被``包裹的也是string类型
    2、golang没有char类型，golang的字符类型是rune，该rune是int的别名，其底层类型就是int
	*/
}
