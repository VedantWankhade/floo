package outbound

import "io"

type AgentRouter interface {
	Route(visitor io.ReadWriteCloser)
	SetConnection(controlConn io.ReadWriteCloser)
}
