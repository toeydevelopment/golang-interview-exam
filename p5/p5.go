package main

import (
	"math/rand"
	"time"
)

func horse(name string) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func run() {

	horse("A")
	horse("B")
	horse("C")
}

func main() {

	run()
}
