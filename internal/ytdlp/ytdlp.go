package ytdlp

import (
	"austere/internal/models"
	"fmt"
	"os/exec"
	"strings"
)

func Download(params *models.BodyParams) error {
	metadata := []string{
		"-metadata Track=\"\"",
		"-metadata Year=\"\"",
		"-metadata Genre=\"\"",
		"-metadata Comment=\"\"",
	}

	if params.Title != "" {
		metadata = append(metadata, fmt.Sprintf("-metadata Title=\"%s\"", params.Title))
	}
	if params.Album != "" {
		metadata = append(metadata, fmt.Sprintf("-metadata Album=\"%s\"", params.Album))
	}
	if params.Artist != "" {
		metadata = append(metadata, fmt.Sprintf("-metadata Artist=\"%s\"", params.Artist))
	}

	cmd := exec.Command("yt-dlp", params.URL, "-x", "--audio-format", "mp3", "--add-metadata",
		"--parse-metadata", "artist:%(artist||uploader)s",
		"--parse-metadata", "title:%(title)s",
		"--parse-metadata", "album:%(album)s",
		"--replace-in-metadata", "album:NA=%(title)s",
		"--ppa", "ThumbnailsConvertor+ffmpeg_o:-c:v png -vf crop='ih'",
		"--ppa", fmt.Sprintf("Metadata:%s", strings.Join(metadata, " ")),
		"--embed-thumbnail",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing yt-dlp:", err)
	}
	fmt.Println(string(output))
	return nil
}
