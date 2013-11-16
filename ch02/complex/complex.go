package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	f := 3.2e5
	display(f)
	x := -7.3 - 8.9i
	display(x)
	y := complex64(-18.3 + 8.9i)
	display(y)
	z := complex(f, 13.2)
	display(z)
	display(real(y))
	display(imag(z))
	display(cmplx.Conj(z))
	display(cmplx.Inf())
	display(cmplx.NaN())
	display(cmplx.Sqrt(z))
}

func display(value interface{}) {
	fmt.Printf("%[1]T: %[1]v\n", value)
}
