package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// agent, err := net.Dial("tcp", ":8090")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Agent connected to server")
	// agent.Write([]byte("register"))
	// var wg sync.WaitGroup
	// wg.Add(1)
	// io.Copy(os.Stdout, agent)
	// wg.Wait()

	agent, err := http.Get("http://localhost:8090/agent/register")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, agent.Body)
}
