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
			fmt.Println("Defer Recover catch.")
			if err := recover(); err != nil {
				fmt.Printf("Catch exception: %v\n", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("Defer Do something")
			}()

			fmt.Println("Do something.")

			if user == "" {
				x := 123
				fmt.Println("Call panic.")
				panic(x)
			}

			fmt.Println("After panic") // NOT execute
		}()

	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Main quit.")
}
