package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func gbecho(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Printf("k: %v - v: %v", k, v)
		log.Println("Hello world!")
	}
	log.Println("Hello world! 2")
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
	/*	logrus.WithField("uri", r.RequestURI).Debug("healthy") */
	fmt.Fprint(w, "OK")
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d) %d\n", i, rand.Intn(25))
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/gb", gbecho)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}
