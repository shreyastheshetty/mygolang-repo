package main

import "fmt"

type mytype int

var x mytype
var y int

func main() {

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 999
	fmt.Println(x)

	y := int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)

}
