package project_delete

import (
	"github.com/google/uuid"
)

type Repository interface {
	ProjectDeleteById(id uuid.UUID) error
}

type UseCase struct {
	repository Repository
}

func New(repository Repository) UseCase {
	return UseCase{
		repository: repository,
	}
}

func (u *UseCase) Call(id uuid.UUID) error {
	err := u.repository.ProjectDeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
