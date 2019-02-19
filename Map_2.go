package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shanken, not strirred`, `martins`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Sc`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunset`},
	}

	// fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
