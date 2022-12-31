package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")
	x := os.Getenv("qqq")
	if x == "" {
		log.Print("nonononon x")
	}

	// now do something with s3 or whatever
	log.Print(s3Bucket)
	log.Print(secretKey)
}
