package main

import (
	"flag"
	"log"

	"github.com/golang-cymru/cardiff-meetup-code/2017-07-grpc/team1/pb/say"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("backend", "127.0.0.1:8888", "Where is the backend?")
	_ = flag.String("message", "Hello World", "Some message to say")

	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()

	_ = say.NewTextToSpeechClient(conn)
}
