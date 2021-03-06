package main

import "fmt"

func ping(pings chan<- string, argc string) {
	pings <- argc
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
