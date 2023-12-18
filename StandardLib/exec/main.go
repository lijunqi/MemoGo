package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func getOutputContinually(name string, args ...string) {
	cmd := exec.Command(name, args...)

	stdoutPipe, _ := cmd.StdoutPipe()
	defer stdoutPipe.Close()

	cmd.Start()
	defer cmd.Wait()

	/*
		p := make([]byte, 1024)
		for {
			stdoutPipe.Read(p)
			fmt.Printf("====> %s\n", string(p))
		}
	*/

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			data := scanner.Text()
			fmt.Printf("===> %s\n", string(data))
		}
	}()

	//if err := cmd.Run(); err != nil {
	//	panic(err)
	//}
}

func main() {
	fmt.Println("=== Start ===")
	getOutputContinually("python", "a.py")
	//getOutputContinually("ping")

	/*
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text()) // Println will add back the final '\n'
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	*/

	/*
		cmd := exec.Command("python", "a.py")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	*/

	/*
		file, err := os.Open("a.log")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		// optionally, resize scanner's capacity for lines over 64K, see next example
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	*/
}
