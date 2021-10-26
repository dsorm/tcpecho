package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil || port > 65535 {
		fmt.Printf("Usage: %v [port]", os.Args[0])
	}

	address := fmt.Sprintf(":%v", port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening for incoming tcp on port 7")
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second*5)
		_, _ = io.Copy(conn, conn)
		conn.Close()
	}

}
