package main

import (
	"log"
	"net"
	"strings"

	"github.com/crackcomm/go-smf/example/demo"
	"github.com/crackcomm/go-smf/rpc"
	flatbuffers "github.com/google/flatbuffers/go"
)

var xreq = strings.Repeat("x", 1000)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.1:20776")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	err = handleClient(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func handleClient(conn net.Conn) (err error) {
	req := buildRequest(xreq)
	hdr := rpc.BuildHeader(1, req, 1792279101)
	log.Printf("Request Header: %s", rpc.NewHeader(hdr))
	if _, err := conn.Write(hdr); err != nil {
		return err
	}
	if _, err := conn.Write(req); err != nil {
		return err
	}
	h, respBuf, err := rpc.ReadRequest(conn)
	if err != nil {
		return err
	}
	log.Printf("Response Header: %s", h)
	resp := demo.GetRootAsResponse(respBuf, 0)
	log.Printf("Response: [ name=%q ]", resp.Name())
	return
}

func buildRequest(s string) []byte {
	builder := flatbuffers.NewBuilder(0)
	name := builder.CreateString(s)
	demo.RequestStart(builder)
	demo.RequestAddName(builder, name)
	resp := demo.RequestEnd(builder)
	builder.Finish(resp)
	return builder.FinishedBytes()
}
