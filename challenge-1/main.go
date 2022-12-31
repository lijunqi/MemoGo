package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 27, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	var wg sync.WaitGroup
	msgList := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}

	wg.Add(3)

	for _, m := range msgList {
		go func(message string, waitgrp *sync.WaitGroup) {
			defer waitgrp.Done()
			updateMessage(message)
			printMessage()
		}(m, &wg)
	}
	wg.Wait()

}
