package main

import (
	"flag"
	pb "github.com/wow-unbelievable/tag-service/proto"
	"github.com/wow-unbelievable/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "p", "7000", "启动端口号")
	flag.Parse()
}

func main()  {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":"+ port)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
