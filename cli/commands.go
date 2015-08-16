package cli

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:    "youtube",
			Aliases: []string{"y"},
			Usage:   "YouTube downloader",
			Action:  doYouTube,
			Flags:   youTubeFlags,
		},
	}
)
