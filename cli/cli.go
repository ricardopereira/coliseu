package cli

import (
	"os"

	"github.com/codegangsta/cli"
)

func Run() {
	app := cli.NewApp()
	app.Name = "Coliseu"
	app.Usage = "Video downloader and audio extractor"
	app.Version = "0.1"
	app.Author = "Ricardo Pereira"
	app.Email = "m@ricardopereira.eu"

	app.Action = func(c *cli.Context) {
		println("Welcome")
	}

	app.Run(os.Args)
}
