package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"test/pb"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	err := check(in.Name)
	if err != nil {
		return nil, err
	}

	caller := getCaller(ctx)
	message := fmt.Sprintf("Hello, caller: %s, name: %s", caller, in.Name)
	header := metadata.Pairs("caller", caller, "message", message)
	grpc.SendHeader(ctx, header)
	return &pb.HelloReply{Message: message}, nil
}

func (s *server) SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	err := check(in.Name)
	if err != nil {
		return nil, err
	}

	caller := getCaller(ctx)
	message := fmt.Sprintf("Hi, caller: %s, name: %s", caller, in.Name)
	header := metadata.Pairs("caller", caller, "message", message)
	grpc.SendHeader(ctx, header)
	return &pb.HelloReply{Message: message}, nil
}

func check(name string) error {
	if len(name) == 0 || strings.EqualFold("error", name) {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid name: %s", name))
	}
	return nil
}

func getCaller(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	caller := md.Get("caller")[0]
	return caller
}

func main() {
	lis, err := net.Listen("tcp", ":16565")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
