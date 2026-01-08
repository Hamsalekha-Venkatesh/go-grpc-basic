# Run these commands to initialize the repo
1. `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

2. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

3. `protoc --go_out=. --go-grpc_out=. proto/main.proto`

5. `go get google.golang.org/grpc`