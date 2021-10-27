package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown() error {

	v := make(chan os.Signal, 1)

	signal.Notify(v, syscall.SIGTERM, syscall.SIGINT)

	<-v

	return nil
}

func doThings() {
	time.Sleep(time.Hour)
}

func main() {
	go gracefulShutdown()
	doThings()
}
