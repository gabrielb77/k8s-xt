package main

import (
	f "fmt"
	"math/rand"
	"os"
	"github.com/sirupsen/logrus"
)

func main() {
	for i := 0; i < 10; i++ {
		f.Printf("%d) %d\n", i, rand.Intn(25))
	}
	setLogger()
}