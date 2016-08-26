package main

import (

	"os"

	"github.com/jawher/mow.cli"
)


func main() {
	app := cli.App("test", "a CLI to do nothing")
	

	app.Action = func() {
		
	}

	app.Run(os.Args)
}
