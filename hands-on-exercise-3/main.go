package main

import "fmt"

var x int = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%T\t%T\t%T", x, y, z)
	fmt.Println(s)

}
