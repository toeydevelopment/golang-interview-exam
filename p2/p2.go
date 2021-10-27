package main

import (
	"fmt"
	"time"
)

var poolSize = 10

func queryData() error {
	fmt.Println("I'm in")
	time.Sleep(time.Second)
	return nil
}

func main() {

	for i := 0; i < 100; i++ {
		go queryData()
	}
}
