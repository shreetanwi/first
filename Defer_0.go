package main

import (
	"fmt"
)

func main() {
	defer foo()
	fmt.Println("Hellow, playground")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}
