package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handlerConn(conn *net.TCPConn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

func main() {
	addrArg := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addrArg)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		go handlerConn(conn)
	}
}
