package main

import (
	"context"
	"log"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/smfrpc/smf-go/example/demo"
	"github.com/smfrpc/smf-go/example/demo_gen"
	"github.com/smfrpc/smf-go/src/smf"
)

type storage struct {
}

func (s *storage) Get(ctx context.Context, req *demo.Request) ([]byte, error) {
	name := req.Name()
	builder := flatbuffers.NewBuilder(len(name))
	offset := builder.CreateByteString(name)
	demo.ResponseStart(builder)
	demo.ResponseAddName(builder, offset)
	resp := demo.ResponseEnd(builder)
	builder.Finish(resp)
	return builder.FinishedBytes(), nil
}

var storageService = demo_gen.NewSmfStorageService(&storage{})

func main() {
	server := new(smf.Server)
	server.RegisterService(storageService)
	log.Println("Starting")
	err := server.ListenAndServe("tcp", ":20766")
	if err != nil {
		log.Fatal(err)
	}
}
