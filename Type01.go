package main

import (
	"fmt"
)

type done int

var x done

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
