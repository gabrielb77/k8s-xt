package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

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
		fmt.Println("id is missing in parameters")
	}
	fmt.Println("id := ", id)
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

	r := mux.NewRouter()
	r.HandleFunc("/gb2/{id}", gbecho2)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/gb", gbecho)
	http.HandleFunc("/health", health)
	/*	http.HandleFunc("/validate-pods", ServeValidatePods) */
	/*	http.HandleFunc("/mutate-pods", ServeMutatePods) */
	http.ListenAndServe(":8080", nil)
}
