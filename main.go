package main

import (
	f "fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func setLogger() {
	logrus.SetLevel(logrus.DebugLevel)

	lev := os.Getenv("LOG_LEVEL")
	if lev != "" {
		llev, err := logrus.ParseLevel(lev)
		if err != nil {
			logrus.Fatalf("cannot set LOG_LEVEL to %q", lev)
		}
		logrus.SetLevel(llev)
	}

	if os.Getenv("LOG_JSON") == "true" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

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

func main() {
	for i := 0; i < 10; i++ {
		f.Printf("%d) %d\n", i, rand.Intn(25))
	}
	/*	setLogger() */
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/gb", gbecho)
	http.ListenAndServe(":8080", nil)
}

/* algoh */
