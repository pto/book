package main

import (
	"fmt"
)

type myInterface interface{}

type myThing struct {
	value int
}

type mySlice []int

func (s mySlice) changeOne(value int) {
	s[0] = value
}

func (s *mySlice) addOne(value int) {
	*s = append(*s, value)
}

type mySliceInterface interface {
	changeOne(int)
}

func main() {
	var x myThing
	fmt.Println("x is", x)
	x.value = 42
	var y myInterface = x
	fmt.Println("x is", x, "y is", y)
	x.value = 123
	fmt.Println("&x is", &x, "&y is", &y)
	fmt.Printf("x is %v (&x is %v/%p), y is %v (&y is %v/%p)\n", x, &x, &x, y, &y, &y)

	var s mySlice
	var r mySliceInterface
	s = append(s, 1)
	r = s
	fmt.Println("s is", s, "r is", r)
	s.changeOne(42)
	fmt.Println("after s.changeOne(42): s is", s, "r is", r)
	r.changeOne(43)
	fmt.Println("after r.changeOne(43): s is", s, "r is", r)
	s.addOne(9)
	s.addOne(10)
	s.addOne(11)
	fmt.Println("after s.addOne(): s is", s, "r is", r)

	s = []int{1, 2, 3}
	r = &s
	s = append(s, 4, 5, 6)
	fmt.Println("s is", s, "r is", r)
	fmt.Printf("s is %v, r is %v\n", s, r)

	var thing1 mySlice = nil
	thing1 = append(thing1, 11, 10, 9)
	thing1.changeOne(1)
	fmt.Println(thing1)

	// var t mySliceInterface
	// t.changeOne(1)

	var i1, i2 myInterface
	t1, t2 := myThing{1}, myThing{2}
	i1, i2 = t1, t2
	fmt.Printf("i1 == i2: %v\n", i1 == i2)
	i2 = t1
	fmt.Printf("after i2 = t1: i1 == i2: %v\n", i1 == i2)
}
