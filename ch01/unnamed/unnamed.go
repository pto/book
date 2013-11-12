package main

import "fmt"

var thing1 struct {
	value int
}

var thing2 struct {
	value int
}

var thing3 struct {
	val int
}

type myThing struct {
	value int
}

var thing4 myThing

type myOtherThing struct {
	value int
}

var thing5 myOtherThing

func main() {
	thing1.value = 123
	thing2 = thing1
	thing3.val = thing1.value
	thing4 = thing1
	thing5 = thing1
	thing5.value = 42
	// thing4 = thing5
	fmt.Println("thing1 is", thing1, "thing2 is", thing2,
		"thing3 is", thing3, "thing4 is", thing4, "thing5 is", thing5)
}
