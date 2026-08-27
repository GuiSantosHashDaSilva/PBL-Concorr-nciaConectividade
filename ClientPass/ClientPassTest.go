package Passageiro

import (
	"fmt"
	"net"
)

var idpass int

func autenticarpass() {
	//var nomepass string

}
func buscarviagem() {
	var origempass string
	var destinopass string
	var dataviagempass string

	fmt.Print("Onde você tá?: ")
	fmt.Scan(&origempass)

	fmt.Print("Onde vc quer ir?: ")
	fmt.Scan(&destinopass)

	fmt.Print("Que dia você quer ir?: ")
	fmt.Scan(&dataviagempass)

	fmt.Printf("Origem: %s Destino:%s", destinopass, origempass)
	fmt.Println()
	fmt.Printf("Dia: %s", dataviagempass)
}

func main() {
	buscarviagem()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("mandando 3 mísseis para sua casa")
	}
	conn.Write([]byte("miau\n"))

}
