package opcua

import (
	"fmt"
	"log/slog"
	"net"
	"sync"
)

type ServerConfig struct {
	Host string
	Port int

	logger *slog.Logger
}

func (s *ServerConfig) addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Server struct {
	config   *ServerConfig
	listener net.Listener

	logger *slog.Logger

	mutex sync.RWMutex
	quit  chan bool
}

func NewServer(config *ServerConfig) *Server {
	server := &Server{
		config: config,
		quit:   make(chan bool),
		logger: config.logger,
	}
	server.logger.Info("server initialized", slog.String("host", config.Host), slog.Int("port", config.Port))
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

	s.mutex.Lock()
	s.listener = listener
	s.mutex.Unlock()

	go s.listenLoop()

	return actualAddr.Port, nil
}

func (s *Server) listenLoop() {
	s.mutex.RLock()
	listener := s.listener
	s.mutex.RUnlock()
	if listener == nil {
		return
	}
	for {
		netConn, err := listener.Accept()
		if err != nil {
			select {
			case <-s.quit:
				return
			default:
				continue
			}
		}
		go func() {
			s.handleConn(netConn)
		}()
	}
}

func (s *Server) handleConn(conn net.Conn) {
}

func (s *Server) Close() error {
	close(s.quit)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.listener == nil {
		return nil
	}
	err := s.listener.Close()
	s.listener = nil
	if err == nil {
		s.logger.Info("server closed successfully")
		return nil
	}
	return fmt.Errorf("failed to close listener: %w", err)
}
