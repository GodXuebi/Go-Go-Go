package main

import "fmt"

func main() {
	messages := make(chan string)
/*默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。这个特性允许我们，
不使用任何其它的同步操作，来在程序结尾等待消息 "ping"*/
	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)
}

/*默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan）通道准备好接收时，
才允许进行发送（chan <-）。*/