package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpcddd/myservice/myservice/pbfile"
	"log"
	"net"
)

var (
	port = flag.Int("port", 4567, "")
)

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("监听异常:", err)
	}
	s := grpc.NewServer()
	pbfile.RegisterMyServiceServer(s, &Server{})
	if err1 := s.Serve(listener); err1 != nil {
		log.Fatal("启动服务异常:", err1)
	}
}

type Server struct {
	pbfile.UnimplementedMyServiceServer
}

func (s *Server) Echo(ctx context.Context, in *pbfile.Message) (*pbfile.Message, error) {
	fmt.Printf("server recv : %+v \n", in)
	return in, nil
}

func (s *Server) EchoSimple(ctx context.Context, in *pbfile.SimpleMessage) (*pbfile.SimpleMessage, error) {
	fmt.Printf("server recv : %+v \n", in)
	return in, nil
}
