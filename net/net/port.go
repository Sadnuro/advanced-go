package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "url para escanear")

func main() {
	// PORT SCANNER
	flag.Parse()

	var wg sync.WaitGroup

	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("El puerto %d esta abierto\n", port)
		}(i)
	}
	wg.Wait()

	// go run net/port.go --site=scanme.webscantest.com
}
