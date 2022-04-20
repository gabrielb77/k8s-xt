package main

import (
	"net/http"
	f "fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		f.Printf("%d) %d\n", i, rand.Intn(25))
	}
}