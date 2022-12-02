package grpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"work_view/internal/adapter/api/grpc/converters"
	"work_view/internal/adapter/api/grpc/pb/work_view"
)

func (s *gRPCExternalServer) ProjectCreate(ctx context.Context, in *v1.ProjectCreateRequest) (*v1.ID, error) {
	log.Printf("Received: %v", in.GetTitle())

	project, err := s.useCases.ProjectCreate.Call(in.GetTitle())
	if err != nil {
		panic(fmt.Sprintf("Project create %v", err))
	}
	return &v1.ID{Id: project.ID.String()}, nil
}

func (s *gRPCExternalServer) ProjectGet(ctx context.Context, in *v1.ID) (*v1.Project, error) {
	pk, err := uuid.Parse(in.GetId())
	if err != nil {
		panic(fmt.Sprintf("Parse UUID %v", err))
	}
	project, err := s.useCases.ProjectGet.Call(pk)
	if err != nil {
		panic(fmt.Sprintf("Project get %v", err))
	}
	pbProject := converters.EntityToPbProject(project)
	return &pbProject, nil
}

func (s *gRPCExternalServer) ProjectDelete(ctx context.Context, in *v1.ID) (*v1.Empty, error) {
	pk, err := uuid.Parse(in.GetId())
	if err != nil {
		panic(fmt.Sprintf("Parse UUID %v", err))
	}
	err = s.useCases.ProjectDelete.Call(pk)
	if err != nil {
		panic(fmt.Sprintf("Project delete %v", err))
	}
	return &v1.Empty{}, nil
}
