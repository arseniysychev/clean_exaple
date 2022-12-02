package some_api

import (
	"github.com/gofrs/uuid"
	"work_view/internal/domain"
)

type SomeApi struct {
	url string
}

func New(url string) (SomeApi, error) {
	repo := SomeApi{
		url: url,
	}
	return repo, nil
}

func (d *SomeApi) ProjectCreate(title string) (domain.Project, error) {
	return domain.Project{}, nil
}

func (d *SomeApi) ProjectGet(id uuid.UUID) (domain.Project, error) {
	return domain.Project{}, nil
}
func (d *SomeApi) ProjectDeleteById(id uuid.UUID) error {
	return nil
}
