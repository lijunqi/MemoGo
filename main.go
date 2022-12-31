package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startingTime := time.Now().UTC()

	fmt.Println("Hello, world!")

	urls := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, url := range urls {
		//checkLink2(url)
		go checkLink(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

	endingTime := time.Now().UTC()
	var duration time.Duration = endingTime.Sub(startingTime)
	fmt.Println("Duration:", duration)
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		return
	}
	fmt.Println(url, "is good")
	c <- "yes"
}

func checkLink2(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		return
	}
	fmt.Println(url, "is good")
}
