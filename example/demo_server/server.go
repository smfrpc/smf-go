package main

import (
	"log"
	"net"

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
	defer conn.Close()
	for {
		if err := handleRequest(conn); err != nil {
			log.Printf("Connection error: %v", err)
			break
		}
	}
}

func handleRequest(conn net.Conn) error {
	hdr, reqBuf, err := rpc.ReadRequest(conn)
	if err != nil {
		return err
	}

	req := demo.GetRootAsRequest(reqBuf, 0)
	// fmt.Printf("Name: %s\n", req.Name())

	resp := buildResponse(req)
	respHdr := rpc.BuildHeader(hdr, resp, 200)

	// debug: print response header
	rhdr := new(rpc.Header)
	rhdr.Init(respHdr, 0)
	log.Printf("Response: %s", rhdr)

	if _, err := conn.Write(respHdr); err != nil {
		return err
	}

	if _, err := conn.Write(resp); err != nil {
		return err
	}

	return nil
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
