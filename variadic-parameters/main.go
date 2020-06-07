package main

import "fmt"

func main() {

	printValues("one", "two", "three")
}

func printValues(values ...string) {

	for i, value := range values {
		fmt.Println(i, ":", value)
	}
}
