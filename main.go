package main

import (
	f "fmt"
	"log"
	"math/rand"
	"net/http"
)

func gbecho(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		f.Printf("k: %v - v: %v", k, v)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	f.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			f.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	/*	logrus.WithField("uri", r.RequestURI).Debug("healthy") */
	log.Prefix("lala ")
	log.Println("health")
	f.Fprint(w, "OK")
}

func main() {
	for i := 0; i < 10; i++ {
		f.Printf("%d) %d\n", i, rand.Intn(25))
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/gb", gbecho)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}

/* algoh */
