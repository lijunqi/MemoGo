package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Start")
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 5)
	scanner.Buffer(buf, 5)
	for scanner.Scan() {
		fmt.Printf("Receive: %s\n", scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		fmt.Println(buf)
	}
}
