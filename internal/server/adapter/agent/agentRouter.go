package agent

import (
	"io"

	"github.com/vedantwankhade/floo/internal/server/core/port/outbound"
)

type TCPAgentRouter struct {
	Conn io.ReadWriteCloser
}

func NewTCPAgentRouter() outbound.AgentRouter {
	return &TCPAgentRouter{}
}

func (r *TCPAgentRouter) Route(visitor io.ReadWriteCloser) {
	io.Copy(r.Conn, visitor)
}

func (r *TCPAgentRouter) SetConnection(controlConn io.ReadWriteCloser) {
	r.Conn = controlConn
}
