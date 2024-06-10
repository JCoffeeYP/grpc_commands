package app

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"wsedge_grpc/config"
	grpcService "wsedge_grpc/internal/controller/grpc"
	"wsedge_grpc/logger"
	pb "wsedge_grpc/proto"
)

func Run(cfg config.Config, logger logger.Logger) {
	fmt.Println("app run")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.StreamInterceptor(grpcService.CustomInterceptor))
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEdgeServiceServer(grpcServer, grpcService.New())

	grpcServer.Serve(lis)
}
