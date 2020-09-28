package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/app"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure/app_logger"
	"os"
)

var application = cli.NewApp()

func main() {
	infrastructure.GetConfiguration()
	logrus, err := app_logger.NewAppLog("", "", "INFO")
	if err != nil {
		log.Errorf("NewAppLog error %+v", err)
	}
	commands(logrus)

}
func commands(logger *log.Logger) {
	application = &cli.App{
		Name: "api",
		Usage: "Run Rest Api",
		Commands: []*cli.Command{
			{
				Name:    "restapi",
				Aliases: []string{"rest","api"},
				Usage:   "complete a task on the list",
				Action:  func(c *cli.Context) error {
					logger.Infof("Run Rest Api")
					app.RegisterGinApp(logger)
					return nil
				},
			},
		},
	}

	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

