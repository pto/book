package main

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	digits := digitsFromCommandLine(1000)
	scaledPi := fmt.Sprint(π(digits))
	fmt.Printf("3.%s\n", scaledPi[1:])
}

func digitsFromCommandLine(default_digits int) int {
	if len(os.Args) > 2 || (len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help")) {
		exitWithUsage()
	}
	if len(os.Args) > 1 {
		if places, err := strconv.Atoi(os.Args[1]); err != nil {
			fmt.Printf("%s: invalid number\n", filepath.Base(os.Args[0]))
			exitWithUsage()
		} else {
			return places
		}
	}
	return default_digits
}

func exitWithUsage() {
	fmt.Printf("usage: %s <number-of-digits>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func π(digits int) *big.Int {
	factor := big.NewInt(int64(digits))
	unity := big.NewInt(0)
	ten := big.NewInt(10)
	exponent := big.NewInt(0)
	unity.Exp(ten, exponent.Add(factor, ten), nil)
	pi := big.NewInt(4)
	left := arccot(big.NewInt(5), unity)
	left.Mul(left, big.NewInt(4))
	right := arccot(big.NewInt(239), unity)
	left.Sub(left, right)
	pi.Mul(pi, left)
	return pi.Div(pi, big.NewInt(0).Exp(ten, ten, nil))
}

func arccot(x, unity *big.Int) *big.Int {
	zero := big.NewInt(0)
	minus_one := big.NewInt(-1)
	x_squared := big.NewInt(0)
	x_squared.Mul(x, x)
	coefficient := big.NewInt(1)
	two := big.NewInt(2)
	divisor := big.NewInt(0)
	divisor.Div(unity, x)
	sum := big.NewInt(0)
	factor := big.NewInt(0)
	factor.Div(divisor, coefficient)
	for factor.Cmp(zero) != 0 {
		sum.Add(sum, factor)
		coefficient.Add(coefficient, two)
		divisor.Mul(divisor, minus_one)
		divisor.Div(divisor, x_squared)
		factor.Div(divisor, coefficient)
	}
	return sum
}
