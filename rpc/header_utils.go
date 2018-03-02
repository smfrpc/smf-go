package rpc

import (
	"io"
)

// ReadHeader - Reads smf RPC header from connection reader.
func ReadHeader(conn io.Reader) (hdr *Header, err error) {
	buf := make([]byte, 16)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return
	}

	hdr = new(Header)
	hdr.Init(buf, 0)
	return
}
