package cli

import (
	"fmt"
	"os"
	"strings"
	"bufio"

	"github.com/bitrise-io/go-utils/colorstring"
	console "github.com/bitrise-io/goinp/goinp"
	"github.com/cheggaaa/pb"
	"github.com/codegangsta/cli"
	"github.com/kennygrant/sanitize"
	youtube "github.com/ricardopereira/coliseu-youtube"
)

var (
	youTubeFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "download, d",
			Value: "",
			Usage: "Download a YouTube video from url or id",
		},
		cli.StringFlag{
			Name:  "file, f",
			Value: "",
			Usage: "Download multiple YouTube videos from a file containing url's or id's",
		},
	}
)

func doYouTube(c *cli.Context) {
	println(colorstring.Redf("YouTube"))
	if downloadInput := c.String("download"); downloadInput != "" {
		println("Download argument:", downloadInput)
		videoDetails(downloadInput, -1)
	} else if fileInput := c.String("file"); fileInput != "" {
		println("File argument:", fileInput)
		readFile(fileInput)
	} else {
		println("Invalid argument")
		return
	}
}

func downloadYouTube(video youtube.Video, option int64) {
	var bar *pb.ProgressBar
	var totalBar int
	var filename = sanitize.BaseName(video.Author) + " - " + sanitize.BaseName(video.Title) + "." + video.GetExtension(int(option))
	// Start download
	video.Download(int(option), filename, func(transferred int, total int) {
		if bar == nil {
			bar = pb.New(total).SetUnits(pb.U_BYTES)
			bar.Start()
			totalBar = total
		}
		bar.Set(transferred)
	})
	// Ended
	if bar != nil {
		bar.Set(totalBar)
		bar.Finish()
	}
}

func videoDetails(input string, option int64) {
	if strings.TrimSpace(input) == "" {
		return
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
		fmt.Println(" - author:", video.Author)
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

	if option != -1 {
		downloadYouTube(video, option)
		fmt.Println(colorstring.Greenf("Done"))
		return
	}

	ask := "Select format to download"
	if option, err := console.AskForInt(ask); err != nil {
		fmt.Println(colorstring.Redf("Error:", err))
	} else if int(option) == cancel {
		fmt.Println(colorstring.Yellowf("Cancelled"))
	} else {
		downloadYouTube(video, option)
		fmt.Println(colorstring.Greenf("Done"))
	}
}

func readFile(file string) {
	if file, err := os.Open(file); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			downloadInput := scanner.Text()
			videoDetails(downloadInput, 0)
		}

		if err = scanner.Err(); err != nil {
			fmt.Println(colorstring.Redf("Error:", err))
		}
	} else {
		fmt.Println(colorstring.Redf("Error:", err))
	}
}
