package main

import (
	"fmt"
	"math/rand"
	"time"
)

func horse(name string) {
	fmt.Printf("%s start", name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("%s done", name)
}

func run() {

	horse("A")
	horse("B")
	horse("C")
}

func main() {

	run()
}
