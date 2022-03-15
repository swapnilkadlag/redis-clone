package main

import (
	"fmt"
	"net"
	"os"
)

const (
	network = "tcp"
	address = "0.0.0.0:6379"
)

func main() {
	listener, error := net.Listen(network, address)
	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}
	defer listener.Close()
	for {
		_, error := listener.Accept()
		if error != nil {
			fmt.Println(error.Error())
			os.Exit(1)
		} else {
			fmt.Println("New connection")
		}
	}
}
