package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/faruqfadhil/bibik/handler"
	"github.com/faruqfadhil/bibik/internal/cli"
	"github.com/faruqfadhil/bibik/internal/repository/cli-repository/document"
)

func initContainer() *handler.CLIHandler {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf(err.Error())
	}
	path := ".bibik"
	cliRepo := document.NewDocument(dirname, path, fmt.Sprintf("%s/%s/%s", dirname, path, "bibik.data"))
	cliService := cli.NewCLIService(cliRepo, spin)
	cliHandler := handler.NewCLIHandler(cliService)
	return cliHandler
}
