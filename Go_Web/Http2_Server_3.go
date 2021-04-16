package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
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
	//当一定时间没请求，则断开连接
	request := make([]byte, 128)
	conn.SetDeadline(time.Now().Add(time.Minute))
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 {
			fmt.Fprintln(os.Stderr, "connection close.")
			break // 客户端关闭了连接
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte("timestamp: " + daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		request = make([]byte, 128)
	}
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
