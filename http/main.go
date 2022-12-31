package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://ra.qtestnet.com")
	if err != nil {
		log.Fatalf("Get err: %v\n", err)
	}
	log.Printf("resp: %v\n", resp)
}
