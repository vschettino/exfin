package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Name:     "Exfin Command Line Interface",
		Usage:    "Housekeeping tasks for Exfin storage.",
		Commands: Commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
