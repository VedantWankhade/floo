package tunnel

import (
	"io"
	"net"
)

type Tunnel struct {
	ClientID string
	Steam    io.ReadWriteCloser
}

func BridgeConnections(reader net.Conn, writer net.Conn) {
	defer reader.Close()
	defer writer.Close()

	io.Copy(reader, writer)
}
