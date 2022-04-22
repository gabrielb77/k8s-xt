package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func gbecho(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "k: %v - v: %v\n", k, v)
		log.Println("Hello world!")
	}
	log.Println("Hello world! 2")
	fmt.Fprintln(w, "Hola 2")
}

func gbecho2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println(w, "id is missing in parameters")
	}
	fmt.Fprintln(w, "id := ", id)
	//call http://localhost:8080/provisions/someId in your browser
	//Output : id := someId
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d) %d\n", i, rand.Intn(25))
	}

	/*	r := mux.NewRouter()
		r.HandleFunc("/gb2/{id}", gbecho2)
		http.HandleFunc("/hello", hello)
		http.HandleFunc("/headers", headers)
		http.HandleFunc("/gb", gbecho)
		http.HandleFunc("/health", health) */
	/*	http.HandleFunc("/validate-pods", ServeValidatePods) */
	/*	http.HandleFunc("/mutate-pods", ServeMutatePods) */
	/*	http.ListenAndServe(":8080", r) */

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
