package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

func main() {
	fmt.Print("=== Start ===")
	//username := "RA-INT\\JLi21"
	//u, err := user.Lookup(username)
	u, err := user.Current()
	if err != nil {
		fmt.Println("xxxError:", err, u.Username)
		return
	}

	fmt.Printf("User: %s (UID: %s)\n", u.Username, u.Uid)

	proc := "C:\\Program Files\\Microsoft Office\\root\\Office16\\EXCEL.EXE"
	cmd := exec.Command(proc)

	//cmd := exec.Command("runas", "/user:JLi21", proc)

	// Set the working directory if needed.
	//cmd.Dir = "path_to_working_directory"

	// Start the command.
	if err := cmd.Start(); err != nil {
		fmt.Printf("xxx Error starting the process: %v\n", err)
		os.Exit(1)
	}

	// Wait for the command to finish.
	if err := cmd.Wait(); err != nil {
		fmt.Printf("xxx Error waiting for the process: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("=== Done ===")
}
