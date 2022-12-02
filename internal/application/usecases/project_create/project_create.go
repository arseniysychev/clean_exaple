package project_create

import (
	"work_view/internal/domain"
)

//type

type Repository interface {
	ProjectCreate(title string) (domain.Project, error)
}

type UseCase struct {
	repository Repository
}

func New(repository Repository) UseCase {
	return UseCase{
		repository: repository,
	}
}

func (u *UseCase) Call(title string) (domain.Project, error) {
	project, err := u.repository.ProjectCreate(title)
	if err != nil {
		return project, err
	}
	return project, nil
}
