package main

import (
	"fmt"
)

func main() {

	x := []int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Print(x)
}
