package rpc

import (
	"fmt"
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

// String - Formats header as debug string.
func (hdr *Header) String() string {
	return fmt.Sprintf("[ compression=%d, bitflags=%d, session=%d, size=%d, checksum=%d, meta=%d ]",
		hdr.Compression(),
		hdr.Bitflags(),
		hdr.Session(),
		hdr.Size(),
		hdr.Checksum(),
		hdr.Meta())
}
