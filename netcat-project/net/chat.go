package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

// Client1 => Server => HandleConnection
// Manipula la conexion para un cliente en especifico
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	message := make(chan string) // Transmision de mensajes del cliente

	go MsgWrite(conn, message)

	//Client1:2560 Platzi.com:2560
	clientName := conn.RemoteAddr().String() // Representacion del puerto de la conexion
	message <- fmt.Sprintf("Welcome to the server, youy name %s\n", clientName)
	messages <- fmt.Sprintf("New client is here, name %s\n", clientName)

	incomingClients <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- message
	messages <- fmt.Sprintf("%s Said bye!\n", clientName)

}

func MsgWrite(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}

func BroadCast() {
	clients := make(map[Client]bool)

	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
		}

	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)

	}
	go BroadCast()
	for { // Acepta la conexion de cada cliente
		conn, err := listener.Accept()
		if err != nil { // No se detiene el programa ya que el error es exclusivo de una conexion
			log.Fatal(err)
			continue
		}
		go HandleConnection(conn)
	}
}
