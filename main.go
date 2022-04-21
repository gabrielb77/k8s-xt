package main

import (
	"bytes"
	"encoding/json"
	f "fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/slackhq/simple-kubernetes-webhook/pkg/admission"
	admissionv1 "k8s.io/api/admission/v1"
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
	f.Fprint(w, "OK")
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info.Println("STD OUT Special Information")
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error.Println("STD ERR")
}

// parseRequest extracts an AdmissionReview from an http.Request if possible
func parseRequest(r http.Request) (*admissionv1.AdmissionReview, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, f.Errorf("Content-Type: %q should be %q",
			r.Header.Get("Content-Type"), "application/json")
	}

	bodybuf := new(bytes.Buffer)
	bodybuf.ReadFrom(r.Body)
	body := bodybuf.Bytes()

	if len(body) == 0 {
		return nil, f.Errorf("admission request body is empty")
	}

	var a admissionv1.AdmissionReview

	if err := json.Unmarshal(body, &a); err != nil {
		return nil, f.Errorf("could not parse admission review request: %v", err)
	}

	if a.Request == nil {
		return nil, f.Errorf("admission review can't be used: Request field is nil")
	}

	return &a, nil
}

func ValidatePods(w http.ResponseWriter, r *http.Request) {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	Info.Println(r.RequestURI)
	Info.Println("received validation request")

	in, err := parseRequest(*r)
	if err != nil {
		Error.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adm := admission.Admitter{
		Logger:  Info,
		Request: in.Request,
	}

	out, err := adm.ValidatePodReview()
	if err != nil {
		e := f.Sprintf("could not generate admission response: %v", err)
		Error.Println(e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jout, err := json.Marshal(out)
	if err != nil {
		e := f.Sprintf("could not parse admission response: %v", err)
		Error.Println(e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	Info.Println("sending response")
	Info.Println("%s", jout)
	f.Fprintf(w, "%s", jout)
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
