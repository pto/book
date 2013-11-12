package main

import (
	"fmt"
	"math"
)

func main() {
	const factor = 3
	i := 20000
	i *= factor
	j := int16(20)
	i += int(j)
	k := uint8(0)
	k = uint8(i)
	fmt.Printf("%T %T %T\n", i, j, k)
	fmt.Println(i, j, k)
	fmt.Println("60020 % 256 is", 60020%256)
	fmt.Println(Uint8FromInt(i))

	var a, b uint8
	a = 0x52
	b = 0x12
	fmt.Printf("0x%x &^ 0x%x is 0x%x\n", a, b, a&^b)
}

func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}
