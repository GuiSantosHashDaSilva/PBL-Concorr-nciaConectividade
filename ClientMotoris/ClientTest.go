package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("mandando 3 mísseis para sua casa")
	}
	conn.Write([]byte("miau\n"))

}
