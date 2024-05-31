package main

import "log"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			log.Println("quit")
			return
		}
	}
}

func CalFibonacci(n int) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			log.Printf("fib[%d] = %d\n", i, <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
