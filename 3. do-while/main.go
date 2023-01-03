package main

import "fmt"

// * Do while loop
func main() {
	i := 3
	for ok := true; ok; ok = i < 3 {
		fmt.Printf("i = %d\n", i)
		i++
	}
}
