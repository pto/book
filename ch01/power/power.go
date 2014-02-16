package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("math.Pow(2,16) is %[1]g (%[1]T)\n", math.Pow(2, 16))
	fmt.Printf("math.MaxUint16 is %[1]d (%[1]T)\n", math.MaxUint16)
	fmt.Printf("1<<16 is %[1]d (%[1]T)\n", 1<<16)
}
