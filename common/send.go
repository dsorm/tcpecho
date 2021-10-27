package common

import (
	"bytes"
	"context"
	"log"
	"net"
	"time"
)

// SendEcho sends a single echo message and returns the response
func SendEcho(data []byte, address string, timeout time.Duration) []byte {
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	c := createConnection(ctx, address)
	defer c.Close()

	return sendAndReceiveData(c, data)
}

// SendEchosAndVerify returns true, nil if all data sent matched data received.
// Otherwise returns false, (data received)
func SendEchosAndVerify(data []byte, echosToSend int, address string, timeout time.Duration) (bool, []byte) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	c := createConnection(ctx, address)
	defer c.Close()

	for i := 0; i < echosToSend; i++ {
		received := sendAndReceiveData(c, data)
		// if not equal
		if bytes.Compare(received, data) != 0 {
			return false, received
		}
	}

	return true, nil
}

func createConnection(ctx context.Context, address string) net.Conn {
	d := net.Dialer{}
	c, err := d.DialContext(ctx, "tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func sendAndReceiveData(c net.Conn, data []byte) []byte {
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
