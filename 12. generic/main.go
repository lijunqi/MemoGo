package main

import "fmt"

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	intRes := add(1, 2)
	fmt.Println("Int Result:", intRes)

	floatRes := add(1.11, 2.22)
	fmt.Println("Float Result:", floatRes)

	strRes := add("abc", "qwer")
	fmt.Println("Str Result:", strRes)
}