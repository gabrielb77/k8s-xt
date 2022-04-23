package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func gbecho(c *gin.Context) {
	/*	for k, v := range req.Header {
		fmt.Fprintf(w, "k: %v - v: %v\n", k, v)
		log.Println("Hello world!")
	} */
	log.Println("Hello world! 2")
	fmt.Println("Hola 2")
}

func gbecho2(c *gin.Context) {
	c.JSON(200, gin.H{"algo": "ok"})
}

func health(c *gin.Context) {
	name := c.Query("name")
	c.String(http.StatusOK, "Hello %s", name)
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

	r.GET("/health", health)
	r.GET("/gb", gbecho)
	r.GET("/gb2", gbecho2)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
