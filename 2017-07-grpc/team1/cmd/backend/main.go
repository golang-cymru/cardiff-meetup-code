package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"

	pb "github.com/golang-cymru/cardiff-meetup-code/2017-07-grpc/team1/pb/say"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8888"
)

type server struct{}

// SaySomething does what it says
func (s *server) SaySomething(ctx context.Context, in *pb.Something) (*pb.Result, error) {
	log.Println("Saying something: " + in.Message)

	file, err := ioutil.TempFile(".", "said")
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
	defer os.Remove(file.Name())

	cmd := exec.Command("say", in.Message, "-o", file.Name())
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name() + ".aiff")

	b, err := ioutil.ReadFile(file.Name() + ".aiff")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Results being sent: " + strconv.Itoa(len(b)))
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
