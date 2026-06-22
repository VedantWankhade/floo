package service

import (
	"io"

	"github.com/vedantwankhade/floo/internal/server/core/port/inbound"
	"github.com/vedantwankhade/floo/internal/server/core/port/outbound"
)

type TunnelServiceImpl struct {
	TunnelRegistry outbound.TunnelRegistry
	AgentRouter    outbound.AgentRouter
}

func NewTunnelService(tunnelRegistry outbound.TunnelRegistry, agentRouter outbound.AgentRouter) inbound.TunnelService {
	return &TunnelServiceImpl{
		TunnelRegistry: tunnelRegistry,
		AgentRouter:    agentRouter,
	}
}

func (s *TunnelServiceImpl) Join(visitor io.ReadWriteCloser) {
	s.AgentRouter.Route(visitor)
}

func (s *TunnelServiceImpl) SetAgentConnection(visitor io.ReadWriteCloser) {
	s.AgentRouter.SetConnection(visitor)
}
