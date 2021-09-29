
https://grpc.io/docs/languages/go/quickstart/

$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$ git clone -b v1.41.0 https://github.com/grpc/grpc-go
$ cd grpc-go/examples/helloworld

$ go run greeter_server/main.go


$ go run greeter_client/main.go

$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

$ go run greeter_server/main.go


$ go run greeter_client/main.go Alice


cd github.com/satya-dillikar/goprojects/grpc-example

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

<- This will generate 2 files helloworld_grpc.pb.go and helloworld.pb.go

cd helloworld
go mod init github.com/satya-dillikar/goprojects/grpc-example/helloworld
go mod tidy

git add helloworld
git commit -m "new helloworld grpc module"
git push origin main

cd greeter_server
go mod init github.com/satya-dillikar/goprojects/grpc-example/greeter_server
go mod tidy
GOPROXY=proxy.golang.org go list -m github.com/satya-dillikar/goprojects/grpc-example/helloworld
go run main.go


cd greeter_client
go mod init github.com/satya-dillikar/goprojects/grpc-example/greeter_client
go mod tidy
GOPROXY=proxy.golang.org go list -m github.com/satya-dillikar/goprojects/grpc-example/helloworld
go run main.go Alice

git add greeter_server greeter_client README.txt
git commit -m "greeter_server greeter_client code to test"
git push origin main