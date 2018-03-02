package rpc

import (
	"io"
	"log"
	"net"
)

// Server - SMF RPC server.
type Server struct {
}

// ListenAndServe - Starts listening on given address and serves connections.
func (server *Server) ListenAndServe(network, address string) (err error) {
	ln, err := net.Listen(network, address)
	if err != nil {
		return
	}
	return server.Serve(ln)
}

// Serve - Starts accepting connections on the listener and serving.
func (server *Server) Serve(ln net.Listener) (err error) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Connection accept error: %v", err)
			continue
		}
		go func() {
			if err := server.Handle(conn); err != nil && err != io.EOF {
				log.Printf("Connection error: %v", err)
			}
		}()
	}
}

// Handle - Handles accepted connection.
func (server *Server) Handle(conn net.Conn) (err error) {
	for {
		// hdr, req, err := ReadRequest(conn)
		// if err != nil {
		// 	return err
		// }

		// Find service by
		// hdr.Meta()

		// Call service.Method(req)

	}

}
