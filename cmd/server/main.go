package main

import (
	"log"

	"github.com/vedantwankhade/floo/internal/server/adapter/agent"
	"github.com/vedantwankhade/floo/internal/server/adapter/api/http"
	"github.com/vedantwankhade/floo/internal/server/core/service"
)

func main() {
	agentRouter := agent.NewTCPAgentRouter()
	tunnelService := service.NewTunnelService(nil, agentRouter)
	// log.Fatal(tcp.NewTCPServer(":8090", tunnelService).Start())
	log.Fatal(http.NewHTTPServer(":8090", tunnelService).Start())

}
