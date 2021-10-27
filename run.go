package main

import (
	"bufio"
	"fmt"
	"github.com/dsorm/tcpecho/common"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil || port > 65535 {
		fmt.Printf("Usage: %v [port]", os.Args[0])
	}

	address := fmt.Sprintf(":%v", port)
	common.ServerStart(address)
	fmt.Printf("Listening for incoming tcp on %v\nPress ENTER to exit\n", address)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
}
