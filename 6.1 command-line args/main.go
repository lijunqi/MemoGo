package main

import (
	"fmt"
	"os"
)

func main() {
	// Method 1
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("i = %d, args = %s\n", i, os.Args[i])
	}

	// Method 2
	for i, arg := range os.Args {
		fmt.Printf("i = %d, args = %s\n", i, arg)
	}
}
