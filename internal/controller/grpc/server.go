package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	pb "wsedge_grpc/proto"
)

type EdgeServer struct {
	pb.UnimplementedEdgeServiceServer
}

func New() *EdgeServer {
	s := &EdgeServer{}
	return s
}

func (s *EdgeServer) RecordCan(stream pb.EdgeService_RecordCanServer) error {
	cnt := 0
	for {
		can, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("%v %d\n", can, cnt)
		cnt += 1
	}
}

func CustomInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	fmt.Println("HELLO WORLD!")
	return nil
}
