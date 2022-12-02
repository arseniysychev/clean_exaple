package models

import (
	"work_view/internal/domain"
)

type Project struct {
	ModelBase
	Title string
}

func (p *Project) MainTo() domain.Project {
	project := domain.Project{
		Title: p.Title,
	}
	project.ID = p.ID
	project.CreatedAt = p.CreatedAt
	project.UpdatedAt = p.UpdatedAt

	return project
}
