// https://go101.org/article/channel-closing.html
package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

/*
 * Don't close (or send values to) closed channels
 *
 * Channel closing principle:
 *     Don't close a channel from the receiver side
 *     Don't close a channel if the channel has multiple concurrent senders.
 * In other words, we should only close a channel in a sender goroutine if the sender is the only sender of the channel.
 */

// ~ 1
type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func CloseChecking() {
	c := make(chan T)
	fmt.Printf("Channel NOT closed. Is closed? %v\n", IsClosed(c)) // false
	close(c)
	select {
	case <-c:
		fmt.Println("Closed!!!!")
	default:
	}
	fmt.Printf("what is: %v\n", <-c)
	fmt.Printf("Channel is closed. Is closed? %v\n", IsClosed(c)) // true
}

// ~ 2
func Bar() {
	bar24x7 := make(chan int, 10) // 此酒吧只能同时招待10个顾客
	for customerId := 0; customerId < 15; customerId++ {
		//time.Sleep(time.Second)
		select {
		case bar24x7 <- customerId: // 试图进入此酒吧
			fmt.Printf("Serve %d\n", customerId)
		default:
			fmt.Print("Customer #", customerId, " leave.\n")
		}
	}
}

// ~ 3
func helper(c chan<- error) {
	time.Sleep(5 * time.Second)
	c <- errors.New("") // send errors/nil on c
}

func NonBlockingRead() error {
	log.Println("do one")

	c := make(chan error, 1)
	go helper(c)

	// * Non-blocking read buffered channel
	// * You can NOT read an unbuffered channel (size 0) without blocking
	quit := false
	for !quit {
		select {
		case <-c:
			log.Println("Quit")
			err := <-c
			log.Printf("resp: %v\n", err)
			quit = true
		default:
			log.Println("Default")
			time.Sleep(time.Second)
		}
	}
	log.Println("do two")

	return nil
}

func main() {
	//CloseChecking()
	//Bar()
	NonBlockingRead()
}
