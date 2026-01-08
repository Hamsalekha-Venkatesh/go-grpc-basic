# Run these commands to initialize the repo
1. `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

2. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

3. ` protoc \
  -I proto \
  --go_out=. \
  --go-grpc_out=. \
  proto/greeter.proto proto/main.proto`

4. `protoc \
  -I proto \
  --go_out=. \
  --go-grpc_out=. \
  proto/greeter.proto proto/stream.proto proto/main.proto proto/add.proto proto/goodbye.proto`

5. `go get google.golang.org/grpc`

`
 protoc \       
  -I proto \
  --go_out=. \
  --go-grpc_out=. \
  proto/stream.proto
`