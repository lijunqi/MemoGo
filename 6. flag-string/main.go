package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "description 1")
	level := flag.String("level", "ERROR", "description 2")

	var flagVar int
	flag.IntVar(&flagVar, "flagname", 1234, "help message for flagname")
	fmt.Printf("1. flagVar: %d\n", flagVar)

	flag.Parse()

	fmt.Printf("2. flagVar: %d\n", flagVar)

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
