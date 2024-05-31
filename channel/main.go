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
	time.Sleep(3 * time.Second)
	c <- errors.New("hello world!!") // send errors/nil on c
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
		case err := <-c:
			log.Println("Quit")
			log.Printf("Received resp: %v\n", err)
			quit = true
		default:
			log.Println("[Default]No message received.")
			time.Sleep(time.Second)
		}
	}
	log.Println("do two")

	return nil
}

// ~ 4
func SelectWithTimeout() {
	log.Println("[Start]Select with Timeout.")
	chan1 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "result"
	}()

	select {
	case res := <-chan1:
		log.Println(res)
	case <-time.After(1 * time.Second):
		log.Println("timeout")
	}
	log.Println("[Done]Select with Timeout.")
}

// ~ 5
func ReceiveFromMultipleChannels() {
	log.Println("[Start]Receive from multi-channel.")
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			log.Println("Received", msg1)
		case msg2 := <-chan2:
			log.Println("Received", msg2)
		}
	}
	log.Println("[Done]Receive from multi-channel.")
}

func main() {
	//CloseChecking()
	//Bar()
	//CalFibonacci(11)

	//NonBlockingRead()
	//SelectWithTimeout()
	ReceiveFromMultipleChannels()
}
