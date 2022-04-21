package main

import (
	f "fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var Info *log.Logger
var Error *log.Logger

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
	log.SetOutput(os.Stderr)
	log.Println("health err")
	log.SetOutput(os.Stdout)
	log.Println("health out")
	f.Fprint(w, "OK")
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info.SetOutput(os.Stdout)
	Info.Println("STD OUT Special Information")
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error.SetOutput(os.Stderr)
	Error.Println("STD ERR")
}

/*
func ValidatePods(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("uri", r.RequestURI)
	logger.Debug("received validation request")

	in, err := parseRequest(*r)
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adm := admission.Admitter{
		Logger:  logger,
		Request: in.Request,
	}

	out, err := adm.ValidatePodReview()
	if err != nil {
		e := fmt.Sprintf("could not generate admission response: %v", err)
		logger.Error(e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jout, err := json.Marshal(out)
	if err != nil {
		e := fmt.Sprintf("could not parse admission response: %v", err)
		logger.Error(e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	logger.Debug("sending response")
	logger.Debugf("%s", jout)
	fmt.Fprintf(w, "%s", jout)
}
*/

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
