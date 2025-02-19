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

// ~ "interface{}" means any type
func printSomething(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("Integer:", val)
	case float64:
		fmt.Println("Float:", val)
	case string:
		fmt.Println("String:", val)
	case abc:
		fmt.Println("Struct abc:", val)
	default:
		fmt.Println("Other type:", val)
	}

	// check value's type
	typedVal, ok := val.(string)
	if ok {
		fmt.Println("typed value is string:", typedVal)
	}
}

func main() {
	num := 1.23
	a := abc{a: "asdf", b: 123, c: true}

	fmt.Println("====== Reflect types ======")
	fmt.Printf("Type Of num: %s\n", reflect.TypeOf(num))
	fmt.Printf("Type Of foo: %s\n", reflect.TypeOf(foo))
	fmt.Printf("Type Of a: %s\n", reflect.TypeOf(a))
	fmt.Printf("Value Of num: %v\n", reflect.ValueOf(num))

	fmt.Println("====== Print something ======")
	printSomething(123)
	printSomething(1.23)
	printSomething("hi")
	printSomething(a)
	printSomething(true)

}
