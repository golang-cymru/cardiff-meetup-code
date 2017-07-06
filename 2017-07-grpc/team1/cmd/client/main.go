package main

import (
	"flag"
	"log"

	"github.com/golang-cymru/cardiff-meetup-code/2017-07-grpc/team1/pb/say"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("backend", "127.0.0.1:8888", "Where is the backend?")
	message := flag.String("message", "Hello World", "Some message to say")

	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
		return
	}

	defer conn.Close()

	client := say.NewTextToSpeechClient(conn)
	ctx := context.Background()

	something := &say.Something{
		Message: *message,
	}

	result, err := client.SaySomething(ctx, something)
	if err != nil {
		log.Fatalf("error returned from say something backend: %s", err)
	}

	// TODO:
	// use bytes from result, store in file
	// run cmd.Exec with afplay
}
