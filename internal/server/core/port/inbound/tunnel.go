package inbound

import "io"

type TunnelService interface {
	Join(vistor io.ReadWriteCloser)
	SetAgentConnection(conn io.ReadWriteCloser)
}
