package opcua

import (
	"fmt"
	"log/slog"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/libgox/buffer"
)

type ServerConfig struct {
	Host string
	Port int

	ReceiverBufferSize int

	ReadTimeout time.Duration

	Logger *slog.Logger

	nextChannelId atomic.Uint32
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

func NewServer(config *ServerConfig) (*Server, error) {
	if config.ReceiverBufferSize == 0 {
		config.ReceiverBufferSize = 1024
	}
	if config.ReceiverBufferSize < 9 {
		return nil, fmt.Errorf("receiver buffer size must be at least 9 bytes")
	}
	server := &Server{
		config: config,
		quit:   make(chan bool),
		logger: config.Logger,
	}
	server.logger.Info("server initialized", slog.String("host", config.Host), slog.Int("port", config.Port))
	return server, nil
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
			s.handleConn(&opcuaConn{
				conn:   netConn,
				buffer: buffer.NewBuffer(s.config.ReceiverBufferSize),
			})
		}()
	}
}

type opcuaConn struct {
	conn   net.Conn
	buffer *buffer.Buffer
}

func (s *Server) handleConn(conn *opcuaConn) {
	channelId := s.nextChannelId()
	channelLogger := s.logger.With(LogRemoteAddr, conn.conn.RemoteAddr().String()).With(LogChannelId, channelId)
	secChannel := newSecureChannel(conn, s.config, channelId, channelLogger)
}

func (s *Server) nextChannelId() uint32 {
	return s.config.nextChannelId.Add(1)
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
