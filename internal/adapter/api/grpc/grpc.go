package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "work_view/internal/adapter/api/grpc/pb/work_view"
	"work_view/internal/application"
	"work_view/internal/conf"
)

func RunServer() {
	cfg := conf.New()
	lis, err := net.Listen("tcp", cfg.GRPC.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	useCases, err := application.NewUseCases(cfg)
	if err != nil {
		log.Fatalf("failed to use cases: %v", err)
	}

	serverGRPC := grpc.NewServer()
	pb.RegisterWorkViewExternalServiceServer(serverGRPC, &gRPCExternalServer{
		useCases: useCases,
	})

	if err := serverGRPC.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("Server success")
}

type gRPCExternalServer struct {
	pb.UnimplementedWorkViewExternalServiceServer
	useCases application.UseCases
}
