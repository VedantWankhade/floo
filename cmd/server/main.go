package main

import (
	"io"
	"log"
	"net"
)

func main() {
	srv, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("error starting the server: %w", err)
	}

	log.Println("Server started listening on :8090")

	agentConn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("error connecting to agent: %w", err)
	}

	log.Println("Connected to agent on :8080")

	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Printf("error receiving request: %w", err)
			continue
		}
		go bridgeConnections(conn, agentConn)
	}
}

func bridgeConnections(visitor, agentConn net.Conn) {
	defer visitor.Close()
	defer agentConn.Close()

	io.Copy(agentConn, visitor)
}
