package enc

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"opcua-go/opcua/uamsg"
	"reflect"
	"sync"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type Decoder interface {
	SetMessageRsvCh(msgCh chan<- *uamsg.Message)
	io.WriteCloser
}

type DecodeContext struct {
	seeker   int64
	opcuaMsg *uamsg.Message
	msgType  uamsg.MessageTypeEnum
	msgSize  uint64
}

type bufferdDecoder struct {
	maxBufferSize int
	mutex         sync.Mutex

	notifyC chan struct{}
	done    chan struct{}

	decodeContext *DecodeContext
	buff          *bytes.Buffer
}

func NewDefaultDecoder(maxBufferSize int) Decoder {

	b := &bufferdDecoder{
		maxBufferSize: 10240000,
		notifyC:       make(chan struct{}, 1),
		buff:          bytes.NewBuffer(nil),
	}

	go b.parseChunks()
	return b
}

func (d *bufferdDecoder) Write(data []byte) (n int, err error) {
	d.mutex.Lock()
	d.buff.Write(data)
	d.mutex.Unlock()
	select {
	case d.notifyC <- struct{}{}:
	default:
	}
	return
}

func (d *bufferdDecoder) Close() error {
	return
}

func (d *bufferdDecoder) SetMessageRsvCh(msgCh chan<- *uamsg.Message) {

}

func (d *bufferdDecoder) parseChunks() {
	for {
		select {
		case <-d.done:
			return
		case _, ok := <-d.notifyC:
			if !ok {
				return
			}
		}
		d.parse()
	}
}

func (d *bufferdDecoder) parse() (interface{}, error) {

}

type superDecoder struct {
	r *bytes.Reader
}

func newSuperDecoder(b []byte) *superDecoder {
	return &superDecoder{
		r: bytes.NewReader(b),
	}
}

func (s *superDecoder) readTo(v interface{}) error {
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Ptr {
		return errors.New("variant is not assignable, must be pointer type")
	}

	switch value.Kind() {
	case reflect.Uint8:
		
		s.r.ReadByte()
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.String:
	case reflect.Slice:
	case reflect.Struct:
	default:

	}
}

func (s *superDecoder) readN(n int) ([]byte, error) {
	if n <= 1 {
		return nil, errors.New("byte`num can`t be less than 1")
	}
	p := make([]byte, n)
	readNum, err := s.r.Read(p)
	if err != nil {
		return nil, err
	}
	if readNum != n {
		return nil, errors.New("no enough bytes")
	}
	return p, nil
}

func (s *superDecoder) readBytes(n int) (byte, error) {
	return s.r.ReadByte()
}