package main

import (
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	pb "github.com/golang-cymru/cardiff-meetup-code/2017-07-grpc/team1/pb/say"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8888"
)

// server is the server
type server struct{}

// SaySomething does what it says
func (s *server) SaySomething(ctx context.Context, in *pb.Something) (*pb.Result, error) {
	log.Println("Saying something: " + in.Message)
	cmd := exec.Command("say", in.Message, "-o", "said.aiff")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("said.aiff")
	if err != nil {
		log.Fatal(err)
	}

	return &pb.Result{Audio: b}, nil
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
