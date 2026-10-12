package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/nir1r7/raft-kv-store/internal/kv"
	"github.com/nir1r7/raft-kv-store/internal/transport"
	"github.com/nir1r7/raft-kv-store/proto/kvpb"
)

func main() {
	// read cmd line
	addr := flag.String("addr", "localhost:5001", "address to listen on")
	flag.Parse()

	// listen at port
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	// register server
	srv := grpc.NewServer()
	kvpb.RegisterKVServer(srv, transport.NewKVServer(kv.NewStore()))

	// start server
	log.Printf("kvnode listening on %s", *addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
