package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	sum := add(x)
	fmt.Print(sum)
}
func add(x []int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
