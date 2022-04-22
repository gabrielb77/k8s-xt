package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func gbecho(c *gin.Context) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "k: %v - v: %v\n", k, v)
		log.Println("Hello world!")
	}
	log.Println("Hello world! 2")
	fmt.Fprintln(w, "Hola 2")
}

func gbecho2(c *gin.Context) {
	c.JSON(200, gin.H{"algo": "ok"})
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d) %d\n", i, rand.Intn(25))
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/gb", gbecho)
	r.GET("/gb2", gbecho2)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
