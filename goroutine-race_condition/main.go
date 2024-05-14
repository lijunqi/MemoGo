package main

import (
	"fmt"
	"time"
)


func subroutine(n *int) {
	for i := 0; i < 100000; i++ {
		*n += 1
	}
}

func main() {
	num := 0

	for i := 0; i < 1; i++ {
		go subroutine(&num)
	}

	for i := 0; i < 100000; i++ {
		num += 1
	}

	time.Sleep(time.Second)

	fmt.Printf("num = %d\n", num)
}