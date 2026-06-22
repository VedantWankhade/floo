package main

import (
	"log"

	"github.com/vedantwankhade/floo/internal/server/adapter/agent"
	"github.com/vedantwankhade/floo/internal/server/adapter/api"
	"github.com/vedantwankhade/floo/internal/server/core/service"
)

func main() {
	agentRouter := agent.NewTCPAgentRouter()
	tunnelService := service.NewTunnelService(nil, agentRouter)
	log.Fatal(api.NewTCPServer(":8090", tunnelService).Start())
}
