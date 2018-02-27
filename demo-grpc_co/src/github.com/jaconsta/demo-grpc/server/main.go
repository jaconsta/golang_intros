package main

import (
  "fmt"
  "log"
  "net"
  "net/http"
  "strings"

  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "golang.org/x/net/context"
  "github.com/jaconsta/demo-grpc/api"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials"
  "google.golang.org/grpc/metadata"
)

// Context keys
type contextKey int

const (
  clientIdKey contextKey = iota
)

func credMatcher(headerName string) (mdName string, ok bool) {
  if headerName == "Login" || headerName == "Password" {
    return headerName, true
  }
  return "", false
}
// authenticateClient check the client credentials
func authenticateClient(ctx context.Context, s *api.Server) (string, error) {
  if md, ok := metadata.FromIncomingContext(ctx); ok {
    clientLogin := strings.Join(md["login"], "")
    clientPassword := strings.Join(md["password"], "")

    if clientLogin != "john" {
      return "", fmt.Errorf("unknown user %s", clientLogin)
    }
    if clientPassword != "doe" {
      return "", fmt.Errorf("bad password %s", clientPassword)
    }

    log.Printf("authenticated client: %s", clientLogin)
    return "42", nil
  }
  return "", fmt.Errorf("missing credentials")
}

// unaryInterceptor calls authenticateClient with currentContext
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
  s, ok := info.Server.(*api.Server)
  if !ok {
    return nil, fmt.Errorf("unable to cast server")
  }
  clientId, err := authenticateClient(ctx, s)
  if err != nil {
    return nil, err
  }
  ctx = context.WithValue(ctx, clientIdKey, clientId)
  return handler(ctx, req)
}

// start a grpc server and wait for connetion
func startGRPCServer(address, certFile, keyFile string) error {  // create a listener
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", address))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  // create server instance
  s := api.Server{}

  // Create TLS credentials
  creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
  if err != nil {
    log.Fatalf("could not load TLS keys: %s", err)
  }

  // grpc options with credentials.
  opts := []grpc.ServerOption{
    grpc.Creds(creds),
    grpc.UnaryInterceptor(unaryInterceptor),
  }
  // create grpc server object
  grpcServer := grpc.NewServer(opts...)

  // Attach ping service to Server
  api.RegisterPingServer(grpcServer, &s)

  // Start Server
  log.Printf("starting HTTP/2 gRPC server on %s", address)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }

  return nil
}

func startRESTServer(address, grpcAddress, certFile, string) error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServerMux(runtime.WithIncomingHeaderMatcher(credMatcher))

  creds, err := credentials.NewClientTLSFromFile(certFile, "")
  if err != nil {
    return fmt.Errorf("could not load TLS certificate: %s", err)
  }

  // Setup the client gRPC options
  opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

  // Register Ping
  err = api.RegisterPingHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
  if err != nil {
    return fmt.Errorf("Could not register service Ping: %s", err)
  }

  log.Printf("starting HTTP/1.1 REST server on %s", address)
  http.ListenAndServe(address, mux)

  return nil
}

// main start a gRPC server and waits for connections
func main() {
  grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)
  restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)
  certFile := "cert/server.crt"
  keyFile := "cert/server.key"

  // fire the gRPC server in a goroutine
  go func() {
    err := startGRPCServer(grpcAddress, certFile, keyFile)
    if err != nil {
      log.Fatalf("failed to start gRPC server: %s", err)
    }
  }()

  // fire the REST server in a goroutine
  go func() {
    err := startRESTServer(restAddress, grpcAddress, certFile)
    if err != nil {
      log.Fatalf("failed to start gRPC server: %s", err)
    }
  }()

  // infinite loop
  log.Printf("Entering infinite loop")
  select {}
}
