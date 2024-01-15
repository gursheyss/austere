package ytdlp

import (
	"austere/internal/models"
	"log"
	"os/exec"
	"strings"
)

func Download(params *models.BodyParams) error {
	metadata := createMetadata(params)

	cmd := exec.Command("yt-dlp", params.URL, "-x", "--audio-format", "mp3", "--add-metadata",
		"--parse-metadata", "artist:%(artist||uploader)s",
		"--parse-metadata", "title:%(title)s",
		"--parse-metadata", "album:%(album)s",
		"--replace-in-metadata", "album:NA=%(title)s",
		"--ppa", "ThumbnailsConvertor+ffmpeg_o:-c:v png -vf crop='ih'",
		"--ppa", "Metadata:"+metadata,
		"--embed-thumbnail",
	)

	output, err := executeCommand(cmd)
	if err != nil {
		log.Println("Error executing yt-dlp:", err)
		return err
	}

	log.Println(string(output))
	return nil
}

func createMetadata(params *models.BodyParams) string {
	metadata := []string{
		"-metadata Track=\"\"",
		"-metadata Year=\"\"",
		"-metadata Genre=\"\"",
		"-metadata Comment=\"\"",
	}

	if params.Title != "" {
		metadata = append(metadata, "-metadata Title=\""+params.Title+"\"")
	}
	if params.Album != "" {
		metadata = append(metadata, "-metadata Album=\""+params.Album+"\"")
	}
	if params.Artist != "" {
		metadata = append(metadata, "-metadata Artist=\""+params.Artist+"\"")
	}

	return strings.Join(metadata, " ")
}

func executeCommand(cmd *exec.Cmd) ([]byte, error) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
