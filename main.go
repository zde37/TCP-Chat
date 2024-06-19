package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("failed to listen on port 7000: %v", err)
	}

	defer listener.Close()
	log.Println("server started on port 7000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}

		go s.newClient(conn)
	}
}
