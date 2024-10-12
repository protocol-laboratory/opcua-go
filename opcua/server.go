package opcua

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/libgox/buffer"

	"github.com/protocol-laboratory/opcua-go/opcua/ua"
)

type ServerConfig struct {
	Host string
	Port int

	ReceiverBufferSize int

	ReadTimeout time.Duration

	Logger *slog.Logger
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
	for {
		if s.config.ReadTimeout > 0 {
			_ = conn.conn.SetReadDeadline(time.Now().Add(s.config.ReadTimeout))
		}
		readLen, err := conn.conn.Read(conn.buffer.WritableSlice())
		if err != nil {
			break
		}
		err = conn.buffer.AdjustWriteCursor(readLen)
		if err != nil {
			break
		}
		for {
			if conn.buffer.ReadableSize() < 8 {
				break
			}

			header := make([]byte, 8)
			err = conn.buffer.PeekExactly(header)
			if err != nil {
				break
			}

			messageLen := int(binary.LittleEndian.Uint32(header[4:8]))

			if conn.buffer.ReadableSize() < messageLen {
				break
			}

			bytes := make([]byte, messageLen)
			err = conn.buffer.ReadExactly(bytes)
			if err != nil {
				break
			}
			dstBytes, err := s.react(conn, bytes)
			if err != nil {
				s.logger.Error("failed to react", slog.String("error", err.Error()))
				break
			}
			write, err := conn.conn.Write(dstBytes)
			if err != nil {
				break
			}
			if write != len(dstBytes) {
				break
			}
			conn.buffer.Compact()
		}
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

func (s *Server) react(conn *opcuaConn, bytes []byte) ([]byte, error) {
	if len(bytes) < 3 {
		return nil, fmt.Errorf("invalid message length")
	}

	messageType := string(bytes[:3])

	s.logger.Info("received message", slog.String("type", messageType))

	var buf *buffer.Buffer
	var err error
	switch messageType {
	case "HEL":
		buf, err = s.handleHello(bytes[8:])
	default:
		return nil, fmt.Errorf("unknown message type: %s", messageType)
	}
	if err != nil {
		return nil, err
	}
	return buf.ReadAll(), nil
}

// handleHello Handle "HEL" (Hello Message)
func (s *Server) handleHello(bytes []byte) (*buffer.Buffer, error) {
	req, err := ua.DecodeMessageHello(buffer.NewBufferFromBytes(bytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode HEL message: %w", err)
	}

	ack := &ua.MessageAcknowledge{
		Version:           req.Version,
		ReceiveBufferSize: req.ReceiveBufferSize,
		SendBufferSize:    req.SendBufferSize,
		MaxMessageSize:    req.MaxMessageSize,
		MaxChunkCount:     64,
	}
	return ack.Buffer()
}
