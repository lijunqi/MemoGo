package main

import (
	"bytes"
	"context"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	log.Println("Start.")
	cmd := exec.CommandContext(ctx, "ping", "127.0.0.1", "-n", "10")
	//cmd := exec.CommandContext(context.Background(), "ping", "127.0.0.1", "-n", "10")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		log.Printf("xxx Start err: %v\n", err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("xxx Wait err: %v\n", err)
	}

	if cmd.Stdout == nil {
		log.Println("Stdout is null")
	} else {
		log.Printf("Stdout: %s\n", cmd.Stdout)
	}

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			log.Println("===> Timeout.")
		case context.Canceled:
			log.Println("===> Canceled.")
		}
	default:
		log.Println("===> Default.")
	}
}
