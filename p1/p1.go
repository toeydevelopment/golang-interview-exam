package main

import "fmt"

var balance = 100

func finalBalance() int {
	return balance
}

func deposit(x int) {
	balance += x
}

func main() {

	n := 1000000

	for i := 0; i < n; i++ {
		go deposit(100)
	}

	fmt.Printf("Balance: %d\n", finalBalance())
}
