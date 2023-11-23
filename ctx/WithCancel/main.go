package main

import (
	"context"
	"log"
	"time"
)

func foo() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					log.Println("Done!")
					time.Sleep(2 * time.Second)
					log.Println("Return!")
					return // returning not to leak the goroutine
				case dst <- n:
					log.Printf("Do %d\n", n)
					time.Sleep(5 * time.Second)
					log.Printf("Complete %d\n", n)
					n++
				}
			}
		}()
		log.Println("******")
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		log.Printf("Get %d\n", n)
		if n == 3 {
			log.Println("Break.")
			break
		}
	}
}

func main() {
	foo()
	log.Println("Out of foo.")
	time.Sleep(10 * time.Second)
	log.Println("Quit.")
}
