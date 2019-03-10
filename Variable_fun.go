package main

import (
	"fmt"
)

var x int
var g func()

func main() {
	f := func() {
		for i := 0; 1 < 3; i++ {
			fmt.Println(i)
		}

	}
	f()
	fmt.Printf("%T\n", f)
	fmt.Println(x)
	fmt.Println("%T\n", x)

	g = f
	g()
	fmt.Printf("this is g %T\n", g)
	fmt.Println("done")
}
