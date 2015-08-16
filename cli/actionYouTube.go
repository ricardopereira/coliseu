package cli

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/ricardopereira/go-get-youtube/youtube"
)

var (
	youTubeFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "download, d",
			Value: "",
			Usage: "Download a YouTube video from url or id",
		},
	}
)

func doYouTube(c *cli.Context) {
	println("YouTube")

	input := c.String("download")

	if input != "" {
		println("Argument:", input)
	}

	// Check if the argument is url or id
	if strings.HasPrefix(input, "http") {
		//"https://www.youtube.com/watch?v={id}"
		//"https://youtube.com/watch?v={id}"
		//"https://youtu.be/{id}"

		// Remove uri
		input = strings.Replace(input, "https://", "", 1)
		input = strings.Replace(input, "http://", "", 1)
		input = strings.Replace(input, "www.", "", 1)
		input = strings.Replace(input, "youtube.com/watch?v=", "", 1)
		input = strings.Replace(input, "youtu.be/", "", 1)

		println("Argument is url")
	} else {
		println("Argument is video id")
	}

	println("Video:", input)

	// Get video interface and metadata
	video, err := youtube.Get(input)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Metadate:")
		fmt.Println(" - title:", video.Title)
		fmt.Println(" - length:", float64(video.Length_seconds)/60.0, "min")
	}
}
