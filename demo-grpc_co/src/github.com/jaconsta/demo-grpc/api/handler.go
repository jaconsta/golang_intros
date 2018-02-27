package api

import (
  "log"

  "golang.org/x/net/context"
)

// represent the gRPC server
type Server struct {

}

// generate a response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
  log.Printf("receive message %s", in.Greeting)
  return &PingMessage{Greeting: "bar"}, nil
}
