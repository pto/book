package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%t %t\n", true, false)
	fmt.Printf("%d %d\n", intForBool(true), intForBool(false))
	t := parseBool("True")
	f := parseBool("FALSE")
	fmt.Printf("%t %t\n", t, f)
}

func intForBool(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func parseBool(s string) bool {
	val, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Printf("cannot parse \"%s\"\n", s)
	}
	return val
}
