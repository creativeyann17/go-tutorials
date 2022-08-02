package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("I'm true")
	}
	m := map[string]int {
		"0": 0,
		"1": 1,
	}
	if pop, ok := m["1"]; ok && returnTrue() {
		fmt.Println(pop)
	}

	switch m["2"] = 2; len(m) {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("don't know")
	}

	switch {
	case len(m) > 0:
		fmt.Println("not empty")
		fallthrough	// force to ignore the implicit break
	case len(m) == 3:
		fmt.Println("And in fact it's 3")
	default:
		fmt.Println("empty")
	}

	// bellow is strange
	var v interface{} = 1
	switch v.(type) {
	case int:
		fmt.Println("int")
		break	// break early
		fmt.Println("int")
	}
}

func returnTrue() bool {
	return true
}