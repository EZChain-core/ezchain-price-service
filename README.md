### Setup
- go version >= 1.14.2
- go protobuf, protoc-gen-go-grpc, go-micro

```bash
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Note: because the new version has replaced so please check and use google protobuf 1.2.0'

```bash
GIT_TAG="v1.2.0" # change as needed
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
```

- Generate proto file:

```bash
make gen
```

- note: in that case that generate error proto, please use v1 via command:

`go get -u github.com/golang/protobuf/protoc-gen-go`

- run server:

```bash
;MICRO_SERVER_ADDRESS=:8000 MICRO_REGISTRY=etcd go run -mod=mod cmd/*.go```
