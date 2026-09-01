package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("fica de xereca ai babão")
		return
	}
	defer ln.Close()
	fmt.Println("Pai ta on")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("fica de xereca ai fofoqueiro")
			return
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	for {

		fmt.Println("Esperando de quebras")

		//aqui bota a lógica q fica no server
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		fmt.Print(message)
	}
}
