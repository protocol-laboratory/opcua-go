package opcua

import (
	"fmt"
	"net"
	"runtime"
	"sync"
	"time"

	"golang.org/x/exp/slog"
)

type ServerConfig struct {
	Host string
	Port int

	Interceptor ServerInterceptor
	Handler     ServerHandler

	ReceiverBufferSize int
	MaxResponseSize    int

	ReadRequestNodeLimit int

	ReadTimeout time.Duration

	Logger *slog.Logger
}

func (s *ServerConfig) addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// Server sort fields by config, opcua biz, inner fields, callbacks, logger, others
type Server struct {
	config *ServerConfig

	channelIdGen   *ChannelIdGen
	sessionManager *SessionManager

	listener net.Listener

	interceptor ServerInterceptor
	handler     ServerHandler

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
	if config.MaxResponseSize <= 0 {
		config.MaxResponseSize = 64 * 1024
	}
	server := &Server{
		config:         config,
		channelIdGen:   &ChannelIdGen{},
		sessionManager: newSessionManager(),
		logger:         config.Logger,
		quit:           make(chan bool),
	}
	if config.Interceptor == nil {
		server.interceptor = &NoopServerInterceptor{}
	} else {
		server.interceptor = config.Interceptor
	}
	if config.Handler == nil {
		server.handler = &NoopServerHandler{}
	} else {
		server.handler = config.Handler
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
			defer func() {
				if e := recover(); e != nil {
					var buf [4096]byte
					n := runtime.Stack(buf[:], false)
					stackInfo := string(buf[:n])
					s.logger.Error("panic, stack: \n"+stackInfo, slog.Any("err", e))
				}
			}()
			s.handleConn(&Conn{
				Conn: netConn,
			})
		}()
	}
}

func (s *Server) handleConn(conn *Conn) {
	s.interceptor.ConnectionOpened(conn)
	channelId := s.channelIdGen.next()
	channelLogger := s.logger.With(LogRemoteAddr, conn.RemoteAddr().String()).With(LogChannelId, channelId)
	channelLogger.Info("starting SecureChannel initialization")
	secChannel := newSecureChannel(conn, s.config, channelId, s.sessionManager, s.interceptor, s.handler, channelLogger)
	err := secChannel.open()
	if err != nil {
		_ = conn.Close()
		s.interceptor.ConnectionClosed(conn)
		channelLogger.Error("failed to open SecureChannel", slog.Any("err", err.Error()))
		return
	}
	err = secChannel.serve()
	if err != nil {
		_ = conn.Close()
		s.interceptor.ConnectionClosed(conn)
		secChannel.logger.Error("processing request error", slog.Any("err", err))
	}
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
