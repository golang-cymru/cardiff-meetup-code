package main

import (
	"log"
	"net"

	pb "github.com/golang-cymru/cardiff-meetup-code/2017-07-grpc/team2/speak"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SaySomething(ctx context.Context, in *pb.SpeakEvent) (*pb.Empty, error) {
	log.Println("In SaySomething")
	return &pb.Empty{}, nil
}

func (s *server) GetVoices(context.Context, *pb.Empty) (*pb.VoiceResponse, error) {
	log.Println("In GetVoices")
	return &pb.VoiceResponse{Voices: []string{"Dummy1", "Dummy2"}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSpeakServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
