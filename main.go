package main

import (
	"flag"
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/implementation"
	"github.com/blackjack200/xyron/xyron"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", 8884))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Print("Listen")
	s := grpc.NewServer()
	t := anticheat.NewSimpleAnticheatServer(implementation.Available)
	xyron.RegisterAnticheatServer(s, t)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
	s.Stop()
}
