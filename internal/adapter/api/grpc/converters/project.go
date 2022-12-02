package converters

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "work_view/internal/adapter/api/grpc/pb/work_view"
	"work_view/internal/domain"
)

func EntityToPbProject(project *domain.Project) pb.Project {
	return pb.Project{
		Id:        project.ID.String(),
		CreatedAt: timestamppb.New(project.CreatedAt),
		UpdatedAt: timestamppb.New(project.UpdatedAt),
		Title:     project.Title,
	}
}
