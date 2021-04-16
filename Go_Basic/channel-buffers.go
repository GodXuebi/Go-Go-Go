package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	var output [2]string

	output[0] = <-messages
	output[1] = <-messages

	fmt.Println(output[0])
	fmt.Println(output[1])
}
