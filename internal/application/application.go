package application

import (
	"work_view/internal/adapter/spi/repository/postgres"
	"work_view/internal/application/usecases/project_create"
	"work_view/internal/application/usecases/project_delete"
	"work_view/internal/application/usecases/project_get"
	"work_view/internal/conf"
)

type UseCases struct {
	ProjectCreate project_create.UseCase
	ProjectGet    project_get.UseCase
	ProjectDelete project_delete.UseCase
}

func NewUseCases(config conf.Config) (useCases UseCases, err error) {
	repo, err := postgres.New(config.DB.DSN())

	useCases = UseCases{
		ProjectCreate: project_create.New(repo),
		ProjectGet:    project_get.New(repo),
		ProjectDelete: project_delete.New(repo),
	}
	return
}
