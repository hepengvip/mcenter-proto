package mcenterproto

import (
	"bufio"
	"io"
)

func ReadFull(r io.Reader, msgSize int) ([]byte, error) {
	buff := make([]byte, msgSize)
	_, err := io.ReadFull(r, buff)
	return buff, err
}

func ReadHeader(r bufio.Reader, sep byte) ([]byte, error) {
	return r.ReadBytes(sep)
}
