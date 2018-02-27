## source

https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2

## commands
$ export GOPATH=/the/directory
$ export GOBIN=$GOPATH/bin
$ export PATH=$PATH:$GOBIN

$ make

or

$ go build -i -v -o bin/server github.com/jaconsta/demo-grpc/server

$ go build -i -v -o bin/client github.com/jaconsta/demo-grpc/client

Both go build is throwing an error of undefined: server: api.RegisterPingServer; client: api.NewPingClient

----

Generate a self signed SSL key.

in github.com/jaconsta/demo-grpc

$ mkdir cert
$ openssl genrsa -out cert/server.key 2048
$ openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
$ openssl req -new -sha256 -key cert/server.key -out cert/server.csr
$ openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650
