package main

import (
	"context"
	"log"
	"net"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/crackcomm/go-smf/example/demo"
	"github.com/crackcomm/go-smf/example/demo_gen"
	"github.com/crackcomm/go-smf/rpc"
)

type demoStorage struct {
}

func (s *demoStorage) Get(ctx context.Context, req *demo.Request) ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)
	name := builder.CreateByteString(req.Name())
	demo.ResponseStart(builder)
	demo.ResponseAddName(builder, name)
	resp := demo.ResponseEnd(builder)
	builder.Finish(resp)
	return builder.FinishedBytes(), nil
}

var storageService = demo_gen.NewStorageService(&demoStorage{})

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

	handle := storageService.MethodHandle(hdr.Meta())
	if handle == nil {
		panic("TODO")
	}

	resp, err := handle(context.TODO(), reqBuf)
	if err != nil {
		panic(err)
	}

	respHdr := rpc.BuildResponseHeader(hdr, resp, 200)

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
