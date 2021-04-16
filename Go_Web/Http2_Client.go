package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	}
	addrArg := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addrArg)
	fmt.Println("Beginning: Want to connect", tcpAddr)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	result := make([]byte, 256)
	_, err = conn.Write([]byte("timestamp"))
	_, err = conn.Read(result)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
