
pb/cat.pb.go: proto/cat.proto
	protoc --go_out=plugins=grpc:pb -Iproto cat.proto
