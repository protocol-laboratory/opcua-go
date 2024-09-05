package opcua

import (
	"fmt"
	"net"
)

type ServerConfig struct {
	Host string
	Port int
}

func (s *ServerConfig) addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Server struct {
	config   *ServerConfig
	listener net.Listener
}

func NewServer(config *ServerConfig) *Server {
	server := &Server{
		config: config,
	}
	return server
}

func (s *Server) Run() (int, error) {
	listener, err := net.Listen("tcp", s.config.addr())
	if err != nil {
		return 0, fmt.Errorf("failed to listen on %s: %w", s.config.addr(), err)
	}

	actualAddr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		return 0, fmt.Errorf("failed to get TCP address from listener")
	}

	if s.config.Port == 0 {
		s.config.Port = actualAddr.Port
	}

	s.listener = listener

	return actualAddr.Port, nil
}

func (s *Server) Close() error {
	if s.listener == nil {
		return nil
	}
	err := s.listener.Close()
	s.listener = nil
	if err == nil {
		return nil
	}
	return fmt.Errorf("failed to close listener: %w", err)
}
