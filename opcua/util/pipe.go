package util

import (
	"bytes"
	"errors"
	"io"
	"sync"
)

// 可以连续写入,写入端可以关闭,可以连续读取
// 可以清空buffer
type BufferedPipe interface {
	io.ReadWriteCloser
	Reset()
}
