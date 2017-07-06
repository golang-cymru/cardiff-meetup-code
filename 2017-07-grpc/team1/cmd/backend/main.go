package main

import (
	"log"
	"net"
	"os/exec"

	pb "github.com/golang-cymru/cardiff-meetup-code/2017-07-grpc/team1/pb/say"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is the server
type server struct{}

// SaySomething does what it says
func (s *server) SaySomething(ctx context.Context, in *pb.Something) (*pb.Result, error) {
	cmd := exec.Command("say", "Hello "+in.Message)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return &pb.Result{Audio: []byte{}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
