build:
	go mod tidy
	go build -v -o service ./cmd/

server:
	go run cmd/main.go server

develop:
	go run cmd/main.go develop

proto_tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	buf mod update api/proto

proto_generate:
	buf generate --template api/proto/buf.gen.yaml api/proto
