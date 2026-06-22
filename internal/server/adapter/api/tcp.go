package api

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/vedantwankhade/floo/internal/server/core/port/inbound"
)

type TCPServer struct {
	Port          string
	TunnelService inbound.TunnelService
}

func NewTCPServer(port string, tunnelService inbound.TunnelService) *TCPServer {
	return &TCPServer{
		Port:          port,
		TunnelService: tunnelService,
	}
}

func (s *TCPServer) Start() error {
	srv, err := net.Listen("tcp", s.Port)
	if err != nil {
		return fmt.Errorf("error starting tcp server: %w", err)
	}
	log.Println("TCP server started on port", s.Port)
	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Println("error accepting connection:", err)
			continue
		}
		go s.TCPHandler(conn)
	}
}

func (s *TCPServer) TCPHandler(conn net.Conn) {
	var in strings.Builder
	io.CopyN(&in, conn, 1)
	if in.String() == "r" {
		s.TunnelService.SetAgentConnection(conn)
		log.Println("Agent registered")
		conn.Write([]byte("you are registered"))
		return
	}
	s.TunnelService.Join(conn)
}
