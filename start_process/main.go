package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd", "/c", "echo", `{"Name": "bob"}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Printf("error decode: %v\n", err)
	}
	if err := cmd.Wait(); err != nil {
		log.Printf("error wait: %v\n", err)
	}
	fmt.Printf("content: %s\n", person.Name)

	err = stdout.Close()
	if err != nil {
		log.Printf("xxx Error: %v\n", err.Error())
	} else {
		log.Println("close good")
	}

	log.Println("Done.")
}
