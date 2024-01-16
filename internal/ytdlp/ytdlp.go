package ytdlp

import (
	"austere/internal/models"
	"log"
	"os/exec"
	"strings"
)

func Download(params *models.BodyParams) error {
	metadata := createMetadata(params)
	outputString := createOutput(params)

	cmd := exec.Command("yt-dlp", params.URL, "-x", "--audio-format", "mp3", "--add-metadata",
		"--parse-metadata", "artist:%(artist||uploader)s",
		"--parse-metadata", "title:%(title)s",
		"--parse-metadata", "album:%(album||title)s",
		"--ppa", "ThumbnailsConvertor+ffmpeg_o:-c:v png -vf crop='ih'",
		"--ppa", "Metadata:"+metadata,
		"-o", outputString,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing yt-dlp: %v", err)
		return err
	}

	log.Println(string(output))
	return nil
}

func createMetadata(params *models.BodyParams) string {
	// the year still doesnt work idk y
	metadata := []string{
		"-metadata Track=\"\"",
		"-metadata Year=\"\"",
		"-metadata Genre=\"\"",
		"-metadata Comment=\"\"",
	}

	if params.Title != "" {
		metadata = append(metadata, "-metadata title=\""+params.Title+"\"")
	}
	if params.Album != "" {
		metadata = append(metadata, "-metadata album=\""+params.Album+"\"")
	}
	if params.Artist != "" {
		metadata = append(metadata, "-metadata artist=\""+params.Artist+"\"")
	}

	return strings.Join(metadata, " ")
}

func createOutput(params *models.BodyParams) string {
	artist := "%(uploader)s"
	album := "%(album)s"
	title := "%(title)s"

	if params.Artist != "" {
		artist = params.Artist
	}
	if params.Album != "" {
		album = params.Album
	}
	if params.Title != "" {
		title = params.Title
	}

	return "./output/" + artist + "/" + album + "/" + title + ".%(ext)s"
}
