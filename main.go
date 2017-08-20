package main

import (
	"os"

	"github.com/jawher/mow.cli"
)

var (
	version string
)

func main() {
	app := cli.App("app", "Description")
	app.Version("v", version)
	app.Run(os.Args)
}
