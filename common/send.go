package common

import (
	"context"
	"log"
	"net"
	"time"
)

func SendEcho(data []byte, address string, timeout time.Duration) []byte {
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	// create the connection
	d := net.Dialer{}
	c, err := d.DialContext(ctx, "tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// send data
	n, err := c.Write(data)
	if n != len(data) {
		log.Fatalf("Error while sending echo: written %v bytes from %v bytes", n, len(data))
	}
	if err != nil {
		log.Fatal(err)
	}

	// receive data
	received := make([]byte, len(data), len(data))
	n, err = c.Read(received)
	if n != len(data) {
		log.Fatalf("Error while sending echo: written %v bytes from %v bytes", n, len(data))
	}
	if err != nil {
		log.Fatal(err)
	}
	return received
}
