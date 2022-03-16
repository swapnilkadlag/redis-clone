package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	const (
		network string = "tcp"
		address string = "0.0.0.0:6379"
	)

	listener, error := net.Listen(network, address)
	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}
	defer listener.Close()

	for {
		connection, error := listener.Accept()
		if error != nil {
			fmt.Println(error.Error())
			os.Exit(1)
		} else {
			go communicate(&connection)
		}
	}
}

func communicate(connection *net.Conn) {
	fmt.Println("New connection")
	reader := bufio.NewReader(*connection)
	//writer := bufio.NewWriter(*connection)
	for {
		reader.ReadByte()
		arr, err := readArray(reader)
		res := process(arr)
		fmt.Println("Response -> " + res)
		//writer.Write([]byte(res))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
