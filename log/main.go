package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f1, _ := os.OpenFile("1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	f2, _ := os.OpenFile("2.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	defaultLogger := log.New(f1, "LOGGER_1: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile)
	mw1 := io.MultiWriter(os.Stdout, f1)
	defaultLogger.SetOutput(mw1)
	defaultLogger.Println("[1]Hello from default logger1.")

	defaultLogger2 := log.New(f2, "", 0)
	mw2 := io.MultiWriter(os.Stdout, f2)
	defaultLogger2.SetOutput(mw2)
	defaultLogger2.Println("[2]Hello from default logger2.")

	log.Println("Normal logger.")

	var f *os.File
	if f == nil {
		log.Println("yes, it's nil", "good")
	}
}
