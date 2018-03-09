package main

import (
	"context"
	"log"
	"strings"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/crackcomm/go-smf/example/demo"
	"github.com/crackcomm/go-smf/example/demo_gen"
	"github.com/crackcomm/go-smf/smf"
)

var xreq = strings.Repeat("x", 1000)

func buildRequest() []byte {
	builder := flatbuffers.NewBuilder(0)
	name := builder.CreateString(xreq)
	demo.RequestStart(builder)
	demo.RequestAddName(builder, name)
	resp := demo.RequestEnd(builder)
	builder.Finish(resp)
	return builder.FinishedBytes()
}

func main() {
	client, err := smf.Dial("tcp", "127.0.0.1:20766")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	store := demo_gen.NewSmfStorageClient(client)
	resp, err := store.Get(context.TODO(), buildRequest())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: [ name=%q ]", resp.Name())
}
