package main

import (
	"fmt"
	"time"
)

/*
 * !!!!!!
 * Recovers only work if the panic happens on the same routine as they were defined.
 * !!!!!!
 * It means that:
 * if a function A calls a goroutine B, and the panic happens inside B, recover wont catch it.
 */
func main() {
	defer fmt.Println("Defer Main") // will this be called when panic?
	var user = ""
	go func() {
		defer func() {
			fmt.Println("Defer caller")
			if err := recover(); err != nil {
				fmt.Printf("Catch exception: %v\n", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("Defer do something")
			}()

			fmt.Println("Do something.")

			if user == "" {
				x := 123
				panic(x)
			}

			fmt.Println("After panic")
		}()

	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main quit.")
}
