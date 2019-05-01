package main

import (
	"log"
	"net"

	pb "github.com/junichi-tanaka/grpc/pb"
	"github.com/junichi-tanaka/grpc/service"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	zapLogger := zap.NewExample()
	defer zapLogger.Sync()

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(zapLogger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
		)),
	)
	catService := &service.MyCatService{}
	// 実行したい実処理をseverに登録する
	pb.RegisterCatServer(server, catService)
	server.Serve(listenPort)
}
