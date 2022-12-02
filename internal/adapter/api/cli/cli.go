package cli

import (
	"log"
	"work_view/internal/application"
	"work_view/internal/conf"
)

func CustomCall() {
	cfg := conf.New()
	useCases, err := application.NewUseCases(cfg)
	if err != nil {
		log.Fatalf("failed to use cases: %v", err)
	}
	project, err := useCases.ProjectCreate.Call("project title")
	if err != nil {
		log.Fatalf("failed to use cases: %v", err)
	}
	log.Println(project)
}
