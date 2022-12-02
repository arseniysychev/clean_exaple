package project_get

import (
	"github.com/google/uuid"
	"work_view/internal/domain"
)

type Repository interface {
	ProjectGet(id uuid.UUID) (domain.Project, error)
}

type UseCase struct {
	repository Repository
}

func New(repository Repository) UseCase {
	return UseCase{
		repository: repository,
	}
}

func (u *UseCase) Call(id uuid.UUID) (domain.Project, error) {
	project, err := u.repository.ProjectGet(id)
	if err != nil {
		return project, err
	}
	return project, nil
}
