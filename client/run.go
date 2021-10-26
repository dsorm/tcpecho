package main

import (
	"fmt"
	"github.com/dsorm/tcpecho/common"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %v [host] [port] [payload]", os.Args[0])
	}

	port, err  := strconv.Atoi(os.Args[2])
	if err != nil || port > 65535 {
		panic("Invalid port number")
	}

	address := fmt.Sprintf("%v:%v", os.Args[1], os.Args[2])
	res := common.SendEcho([]byte(os.Args[3]), address, 10 * time.Second)

	fmt.Printf("[%v] %v\n", address, string(res))

}
