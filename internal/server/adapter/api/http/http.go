package http

import (
	"log"
	"net/http"

	"github.com/vedantwankhade/floo/internal/server/core/port/inbound"
)

type HTTPServer struct {
	Port          string
	TunnelService inbound.TunnelService
}

func NewHTTPServer(port string, tunnelService inbound.TunnelService) *HTTPServer {
	return &HTTPServer{
		Port:          port,
		TunnelService: tunnelService,
	}
}

func (s *HTTPServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/agent/register", s.RegisterAgentHandler)
	mux.HandleFunc("/", s.VisitorHandler)
	log.Println("HTTP server started on port", s.Port)
	return http.ListenAndServe(s.Port, mux)
}

func (s *HTTPServer) RegisterAgentHandler(w http.ResponseWriter, r *http.Request) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "failed to hijack agent control connection", http.StatusInternalServerError)
		return
	}
	conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, "failed to hijack agent control connection: "+err.Error(), http.StatusInternalServerError)
		return
	}
	s.TunnelService.SetAgentConnection(conn)
	log.Println("Agent registered")
	// msg := "You are hijacked"
	// fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\nContent-Length: %d\r\n\r\n%s", len(msg)+99, msg)
}

func (s *HTTPServer) VisitorHandler(w http.ResponseWriter, r *http.Request) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "failed to hijack agent control connection", http.StatusInternalServerError)
		return
	}
	conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, "failed to hijack agent control connection: "+err.Error(), http.StatusInternalServerError)
		return
	}
	s.TunnelService.Join(conn)
}
