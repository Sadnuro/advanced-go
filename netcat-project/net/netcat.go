package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("p", 3090, "puerto")
	host = flag.String("h", "localhost", "host")
)

func main() {
	// netcat -> host:port
	// Escribir -> host:port
	// Leer -> host:port
	// >hola -> host:port -> other conexion: hola

	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // conn como lector
		done <- struct{}{}
	}()
	CopyContent(conn, os.Stdin) // conn como escritor
	conn.Close()
	<-done
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
