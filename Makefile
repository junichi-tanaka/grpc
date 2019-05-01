
.PHONY: build
build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/client cmd/client/main.go

pb/cat.pb.go: proto/cat.proto
	protoc --go_out=plugins=grpc:pb -Iproto cat.proto
