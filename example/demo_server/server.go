package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/cespare/xxhash"
	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/crackcomm/go-smf/example/demo"
	"github.com/crackcomm/go-smf/rpc"
)

func main() {
	ln, err := net.Listen("tcp", ":20766")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Listening")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	hdr, err := readHeader(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Size: %v\n", hdr.Size())
	fmt.Printf("Meta: %v\n", hdr.Meta())
	fmt.Printf("Session: %v\n", hdr.Session())
	fmt.Printf("Checksum: %v\n", hdr.Checksum())
	fmt.Printf("Bitflags: %v\n", hdr.Bitflags())
	fmt.Printf("Compression: %v\n\n", hdr.Compression())

	reqBuf := make([]byte, hdr.Size())
	if _, err := io.ReadFull(conn, reqBuf); err != nil {
		log.Fatal(err)
	}

	req := demo.GetRootAsRequest(reqBuf, 0)
	// fmt.Printf("Name: %s\n", req.Name())

	resp := buildResponse(req)
	fmt.Printf("Response size: u%d\n\n", uint32(len(resp)))
	respHdr := buildHeader(hdr, resp)
	fmt.Printf("respHdr size: u%d\n\n", uint32(len(respHdr)))

	// debug: print response header
	rhdr := new(rpc.Header)
	rhdr.Init(respHdr, 0)
	fmt.Printf("Size: %v\n", rhdr.Size())
	fmt.Printf("Meta: %v\n", rhdr.Meta())
	fmt.Printf("Session: %v\n", rhdr.Session())
	fmt.Printf("Checksum: %v\n", rhdr.Checksum())
	fmt.Printf("Bitflags: %v\n", rhdr.Bitflags())
	fmt.Printf("Compression: %v\n", rhdr.Compression())

	if _, err := conn.Write(respHdr); err != nil {
		log.Fatal(err)
	}

	if _, err := conn.Write(resp); err != nil {
		log.Fatal(err)
	}
}

// readHeader - Reads smf RPC header from connection.
func readHeader(conn io.Reader) (hdr *rpc.Header, err error) {
	buf := make([]byte, 16)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return
	}

	hdr = new(rpc.Header)
	hdr.Init(buf, 0)
	return
}

// buildHeader - Builds response header from response bytes and request header.
func buildHeader(req *rpc.Header, resp []byte) []byte {
	builder := flatbuffers.NewBuilder(20) // [1]
	res := rpc.CreateHeader(builder,
		0,                                         // compression int8,
		0,                                         // bitflags int8,
		req.Session(),                             // session uint16,
		uint32(len(resp)),                         // size uint32,
		uint32(math.MaxUint32&xxhash.Sum64(resp)), // checksum uint32,
		0, // req.Meta(), //	meta uint32
	)
	fmt.Printf("rpc.CreateHeader: u%d\n\n", res)
	builder.Finish(res)
	// TODO(crackcomm): builder prepends 4 bytes
	// the header is the last 16 bytes of message
	// so I did [^1] 20 bytes allocation anyway
	// I have no idea why it does that
	return builder.FinishedBytes()[4:]
}

func buildResponse(req *demo.Request) []byte {
	builder := flatbuffers.NewBuilder(0)
	name := builder.CreateByteString(req.Name())
	demo.ResponseStart(builder)
	demo.ResponseAddName(builder, name)
	resp := demo.ResponseEnd(builder)
	builder.Finish(resp)
	return builder.FinishedBytes()
}
