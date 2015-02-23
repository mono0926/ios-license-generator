package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ios-license-generator"
	app.Version = Version
	app.Usage = ""
	app.Author = "Masayuki Ono"
	app.Email = "mono0926@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
