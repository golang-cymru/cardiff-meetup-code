package main

import (
    "flag"
    "log"

    "../../pb/say"

    "google.golang.org/grpc"
)


func main() {
    backend := flag.String("backend", "127.0.0.1:8888", "Where is the backend?")
    message := flag.String("message", "Hello World", "Some message to say")

    flags.Parse()

    conn, err := grpc.Dial(*backend, grpc.WithInsecure())
    if (err != nil) {
        log.Panic(err)
    }



    defer conn.Close()
}
