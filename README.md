# ms Work view
Service works with dynamic entities

### Requirements

- Docker v20.10.21+
- Docker-Compose v1.29.2+
- Go v1.19+

export GOPATH=~/projects/go
export PATH=$PATH:"$GOPATH"/bin


go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"
