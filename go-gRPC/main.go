package main

import (
	"log"
	pb "my-gRPC/hello"
)

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\hello\student.proto

func main() {
	s := &pb.Student{
		Name:   "Jacky",
		Age:    123,
		Gender: "Male",
		Number: 456,
	}

	log.Println(
		s.GetName(),
		s.GetAge(),
		s.GetGender(),
		s.GetNumber(),
	)
}
