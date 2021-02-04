package main

import (
	"context"
	"fmt"
	"net"
	"proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var queue []string

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func enqueue(queue []string, element string) []string {
	queue = append(queue, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []string) []string {
	element := queue[0] // The first element is the one to be dequeued.
	fmt.Println("Dequeued:", element)
	return queue[1:] // Slice off the element once it is dequeued.
}

func (s *server) Message(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	var a string = request.GetA()

	var result string = a
	// Make a queue of ints.

	queue = enqueue(queue, result)

	fmt.Println("Queue:", queue)

	//queue = dequeue(queue)

	return &proto.Response{Result: result}, nil
}
