package rpc

import (
	"fmt"
	"io"
	"math"

	"github.com/cespare/xxhash"

	flatbuffers "github.com/google/flatbuffers/go"
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

// ReadRequest - Reads request header and body.
func ReadRequest(conn io.Reader) (hdr *Header, req []byte, err error) {
	hdr, err = ReadHeader(conn)
	if err != nil {
		return
	}
	req = make([]byte, hdr.Size())
	_, err = io.ReadFull(conn, req)
	return
}

// BuildHeader - Builds response header from response bytes and request header.
func BuildHeader(req *Header, resp []byte, status uint32) []byte {
	builder := flatbuffers.NewBuilder(20) // [1]
	res := CreateHeader(builder,
		0,                                         // compression int8,
		0,                                         // bitflags int8,
		req.Session(),                             // session uint16,
		uint32(len(resp)),                         // size uint32,
		uint32(math.MaxUint32&xxhash.Sum64(resp)), // checksum uint32,
		status, //	meta uint32
	)
	builder.Finish(res)
	// TODO(crackcomm): builder prepends 4 bytes
	// the header is the last 16 bytes of message
	// so I did [^1] 20 bytes allocation anyway
	// I have no idea why it does that
	return builder.FinishedBytes()[4:]
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
