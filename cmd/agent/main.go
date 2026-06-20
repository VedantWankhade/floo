package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	agentSrv, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error starting agent: %w", err)
	}

	log.Println("Agent started listening on :8080")

	for {
		conn, err := agentSrv.Accept()
		if err != nil {
			log.Printf("error reqceing request: %w", err)
			continue
		}
		io.Copy(os.Stdout, conn)
	}
}
