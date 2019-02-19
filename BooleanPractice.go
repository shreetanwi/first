package main

import (
	"fmt"
)

var x = 20
var y = 10

func main() {
	fmt.Println(x)
	fmt.Println(y)

	if x == y {
		fmt.Println("x is equal to y")
	} else if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is less than y")
	}

	// String comparision
	a := "Ashish"
	b := "Sinha"
	c := "Ashish"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	if a == b {
		fmt.Println("a is equal to b")
	}
	if a == c {
		fmt.Println("a is equal to c")
	}

}
