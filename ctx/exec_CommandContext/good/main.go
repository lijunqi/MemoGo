package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	log.Println("Start.")
	if err := exec.CommandContext(ctx, "cmd.exe", "/k", "..\\a.bat").Run(); err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Println("Success.")
	}

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			log.Println("Timeout.")
		case context.Canceled:
			log.Println("Canceled.")
		}
	default:
		log.Println("Done.")
	}
}
