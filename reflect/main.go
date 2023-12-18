package main

import (
	"fmt"
	"reflect"
)

type abc struct {
	a string
	b int
	c bool
}

func foo() {
	fmt.Println("This is foo.")
}

func main() {
	num := 1.23
	a := abc{}

	fmt.Printf("Type num: %s\n", reflect.TypeOf(num))
	fmt.Printf("Type foo: %s\n", reflect.TypeOf(foo))
	fmt.Printf("Type a: %s\n", reflect.TypeOf(a))
	fmt.Printf("Value: %f\n", reflect.ValueOf(num))
}
