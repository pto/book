package main

import "fmt"

func main() {
	fmt.Println(thing(1))
}

func thing(i int) string {
	switch i {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	default:
		return "other"
	}
}
