package opcua

import (
	"context"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"github.com/shoothzj/gox/buffer"
	"github.com/shoothzj/gox/netx"
	"log/slog"
	"net"
	"opcua-go/opcua/ua"
	"sync"
)

type ClientConfig struct {
	Address          netx.Address
	BufferMax        int
	SendQueueSize    int
	PendingQueueSize int
	TlsConfig        *tls.Config

	Logger *slog.Logger
}

type sendRequest struct {
	buf      *buffer.Buffer
	callback func(*buffer.Buffer, error)
}

type Client struct {
	config *ClientConfig
	logger *slog.Logger

	conn         net.Conn
	eventsChan   chan *sendRequest
	pendingQueue chan *sendRequest
	buffer       *buffer.Buffer
	ctx          context.Context
	ctxCancel    context.CancelFunc
}

func (c *Client) Hello(message *ua.MessageHello) (*ua.MessageAcknowledge, error) {
	buf, err := message.Buffer()
	if err != nil {
		return nil, err
	}
	bufResp, err := c.Send(buf)
	if err != nil {
		return nil, err
	}
	return ua.DecodeMessageAcknowledge(bufResp)
}

func (c *Client) Send(buf *buffer.Buffer) (*buffer.Buffer, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var result *buffer.Buffer
	var err error
	c.sendAsync(buf, func(resp *buffer.Buffer, e error) {
		result = resp
		err = e
		wg.Done()
	})
	wg.Wait()
	if err != nil {
		return nil, err
	}
	err = result.Skip(8)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) sendAsync(buf *buffer.Buffer, callback func(*buffer.Buffer, error)) {
	select {
	case <-c.ctx.Done():
		callback(nil, ErrClientClosed)
	default:
		sr := &sendRequest{
			buf:      buf,
			callback: callback,
		}
		c.eventsChan <- sr
	}
}

func (c *Client) read() {
	for req := range c.pendingQueue {
		n, err := c.conn.Read(c.buffer.WritableSlice())
		if err != nil {
			req.callback(nil, err)
			c.close()
			break
		}
		err = c.buffer.AdjustWriteCursor(n)
		if err != nil {
			req.callback(nil, err)
			c.close()
			break
		}
		if c.buffer.ReadableSize() < 8 {
			continue
		}
		bytes := make([]byte, 8)
		err = c.buffer.PeekExactly(bytes)
		if err != nil {
			req.callback(nil, err)
			c.close()
			break
		}
		length := int(binary.LittleEndian.Uint32(bytes[4:8]))
		if c.buffer.ReadableSize() < length {
			continue
		}
		// in case ddos attack
		if length > c.buffer.Capacity() {
			req.callback(nil, fmt.Errorf("response length %d is too large", length))
			c.close()
			break
		}
		data := make([]byte, length)
		err = c.buffer.ReadExactly(data)
		if err != nil {
			req.callback(nil, err)
			c.close()
			break
		}
		c.buffer.Compact()
		req.callback(buffer.NewBufferFromBytes(data), nil)
	}
}

func (c *Client) write() {
	for req := range c.eventsChan {
		bytes, err := req.buf.ReadNBytes(req.buf.ReadableSize())
		if err != nil {
			req.callback(nil, err)
			c.close()
			break
		}
		n, err := c.conn.Write(bytes)
		if err != nil {
			req.callback(nil, err)
			c.close()
			break
		}
		if n != len(bytes) {
			req.callback(nil, fmt.Errorf("write %d bytes, but expect %d bytes", n, len(bytes)))
			c.close()
			break
		}
		c.pendingQueue <- req
	}
}

func (c *Client) close() {
	c.Close()
}

func (c *Client) Close() {
	c.ctxCancel()
	_ = c.conn.Close()
	close(c.eventsChan)
	close(c.pendingQueue)
}

func NewClient(config *ClientConfig) (*Client, error) {
	conn, err := netx.Dial(config.Address, config.TlsConfig)

	if err != nil {
		return nil, err
	}
	if config.SendQueueSize == 0 {
		config.SendQueueSize = 1000
	}
	if config.PendingQueueSize == 0 {
		config.PendingQueueSize = 1000
	}
	if config.BufferMax == 0 {
		config.BufferMax = 512 * 1024
	}

	ctx, cancel := context.WithCancel(context.Background())

	client := &Client{
		config: config,
		logger: config.Logger,

		conn:         conn,
		eventsChan:   make(chan *sendRequest, config.SendQueueSize),
		pendingQueue: make(chan *sendRequest, config.PendingQueueSize),
		buffer:       buffer.NewBuffer(config.BufferMax),
		ctx:          ctx,
		ctxCancel:    cancel,
	}
	go func() {
		client.read()
	}()
	go func() {
		client.write()
	}()
	return client, nil
}
