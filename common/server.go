package common

import (
	"fmt"
	"io"
	"net"
)

func ServerStart(address string) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 64; i++ {
		var conn net.Conn
		var err error
		go func() {
			for {
				conn, err = l.Accept()
				if err != nil {
					fmt.Println(err)
				}
				_, _ = io.Copy(conn, conn)
				_ = conn.Close()
			}
		}()
	}

}
