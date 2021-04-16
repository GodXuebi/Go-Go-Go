package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	done <- true
}
//worker waits for main
func main() {
	done := make(chan bool, 1)
	go worker(done)
	time.Sleep(time.Second)
	fmt.Println("done")
	<-done
}
