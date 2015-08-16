package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/colorstring"
	"github.com/bitrise-io/goinp/goinp"
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
	println(colorstring.Redf("YouTube"))

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

	cancel := 0
	// Get video interface and metadata
	video, err := youtube.Get(input)

	if err != nil {
		fmt.Println(colorstring.Redf("Error:", err))
		os.Exit(0)
	} else {
		fmt.Println("Metadata:")
		fmt.Println(" - title:", video.Title)
		fmt.Println(" - length:", float64(video.Length_seconds)/60.0, "min")

		fmt.Println(" - format:")
		i := 0
		for i < len(video.Formats) {
			format := video.Formats[i]
			fmt.Println("    ", i, "-", format.Quality, format.Video_type)
			i++
		}
		cancel = i
		fmt.Println("    ", cancel, "-", "Cancel")
	}

	ask := "Select format to download"
	if option, err := goinp.AskForInt(ask); err != nil {
		fmt.Println(colorstring.Redf("Error:", err))
	} else if int(option) == cancel {
		fmt.Println(colorstring.Yellowf("Cancelled"))
	} else {
		video.Download(int(option), video.Id+"."+video.GetExtension(int(option)))
		fmt.Println(colorstring.Greenf("Success"))
	}

}
