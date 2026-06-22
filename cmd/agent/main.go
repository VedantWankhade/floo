package main

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	agent, err := net.Dial("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Agent connected to server")
	agent.Write([]byte("register"))
	var wg sync.WaitGroup
	wg.Add(1)
	io.Copy(os.Stdout, agent)
	wg.Wait()
}
