package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"work_view/internal/adapter/api/grpc"
	"work_view/internal/adapter/spi/repository/postgres"
	"work_view/internal/develop"
)

func commandMigrate(cCtx *cli.Context) error {
	postgres.Migrate()
	return nil
}

func commandRunServer(cCtx *cli.Context) error {
	grpc.RunServer()
	return nil
}

func commandDevelop(cCtx *cli.Context) error {
	develop.Develop()
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Work View"
	app.Usage = "backend microservice"
	app.Commands = []*cli.Command{{
		Name:  "migrate",
		Usage: "Migrate DB",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "up",
				Usage: "Apply migrations",
			},
		},
		Action: commandMigrate,
	}, {
		Name:   "server",
		Usage:  "Run server",
		Action: commandRunServer,
	}, {
		Name:   "develop",
		Usage:  "Develop test",
		Action: commandDevelop,
	},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
