package main

import (
	"fmt"
	"webserver/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "Mcmillan",
	}
	fmt.Println(u)
}
