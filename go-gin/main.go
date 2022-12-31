package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping/:delay", ping)

	r.Run(":8000")
}

func ping(c *gin.Context) {
	log.Println("Receive.")
	delay := c.Param("delay")
	log.Println("Delay: ", delay)
	cnt, _ := strconv.Atoi(delay)
	time.Sleep(time.Duration(cnt) * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
