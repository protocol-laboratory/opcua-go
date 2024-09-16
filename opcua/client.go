package opcua

import (
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"github.com/shoothzj/gox/buffer"
	"github.com/shoothzj/gox/netx"
	"log/slog"
	"net"
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
	closeCh      chan struct{}
}

func (c *Client) Hello(message *MessageHello) (*MessageAcknowledge, error) {
	buf, err := message.Buffer()
	if err != nil {
		return nil, err
	}
	bufResp, err := c.Send(buf)
	if err != nil {
		return nil, err
	}
	return DecodeMessageAcknowledge(bufResp)
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
	return result, nil
}

func (c *Client) sendAsync(buf *buffer.Buffer, callback func(*buffer.Buffer, error)) {
	sr := &sendRequest{
		buf:      buf,
		callback: callback,
	}
	c.eventsChan <- sr
}

func (c *Client) read() {
	for {
		select {
		case req := <-c.pendingQueue:
			n, err := c.conn.Read(c.buffer.WritableSlice())
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			err = c.buffer.AdjustWriteCursor(n)
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			if c.buffer.ReadableSize() < 8 {
				continue
			}
			bytes := make([]byte, 8)
			err = c.buffer.PeekExactly(bytes)
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			length := int(binary.LittleEndian.Uint32(bytes[4:8]))
			if c.buffer.ReadableSize() < length {
				continue
			}
			// in case ddos attack
			if length > c.buffer.Capacity() {
				req.callback(nil, fmt.Errorf("response length %d is too large", length))
				c.closeCh <- struct{}{}
				break
			}
			data := make([]byte, length)
			err = c.buffer.ReadExactly(data)
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			c.buffer.Compact()
			req.callback(buffer.NewBufferFromBytes(data), nil)
		case <-c.closeCh:
			return
		}
	}
}

func (c *Client) write() {
	for {
		select {
		case req := <-c.eventsChan:
			bytes, err := req.buf.ReadNBytes(req.buf.ReadableSize())
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			n, err := c.conn.Write(bytes)
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			if n != len(bytes) {
				req.callback(nil, fmt.Errorf("write %d bytes, but expect %d bytes", n, len(bytes)))
				c.closeCh <- struct{}{}
				break
			}
			c.pendingQueue <- req
		case <-c.closeCh:
			return
		}
	}
}

func (c *Client) Close() {
	_ = c.conn.Close()
	c.closeCh <- struct{}{}
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

	client := &Client{
		config: config,
		logger: config.Logger,

		conn:         conn,
		eventsChan:   make(chan *sendRequest, config.SendQueueSize),
		pendingQueue: make(chan *sendRequest, config.PendingQueueSize),
		buffer:       buffer.NewBuffer(config.BufferMax),
		closeCh:      make(chan struct{}),
	}
	go func() {
		client.read()
	}()
	go func() {
		client.write()
	}()
	return client, nil
}
