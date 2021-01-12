package main

import (
	"bytes"
	"io"
	"net"
)

func readNetStreamSz(w io.Writer, r io.Reader, receivedBytes int64, bufferSize int64, contentLength int64) error {
	for (contentLength - receivedBytes) > bufferSize {
		_, err := io.CopyN(w, r, bufferSize)
		if err != nil {
			return err
		}
		receivedBytes += bufferSize
	}
	return nil
}

func readRequestBody(conn net.Conn, length int64) []byte {
	var b bytes.Buffer
	var receivedBytes int64
	readNetStreamSz(&b, conn, receivedBytes, 4096, length)
	return b.Bytes()
}
