package ytdlp

import (
	"austere/internal/krakenfiles"
	"austere/internal/models"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Download(params *models.BodyParams) error {
    metadata := createMetadata(params)
    outputString := createOutput(params)
    url := params.URL

    if strings.Contains(url, "krakenfiles.com") {
        var err error
        url, err = krakenfiles.ExtractDownloadURL(url)
        if err != nil {
            return err
        }
    }

    cmd := exec.Command("yt-dlp", url, "-x", "--audio-format", "mp3", "--add-metadata",
        "--parse-metadata", "artist:%(artist||uploader)s",
        "--parse-metadata", "title:%(title)s",
        "--parse-metadata", "album:%(album||title)s",
        "--embed-thumbnail",
        "--ppa", "ThumbnailsConvertor+ffmpeg_o:-c:v png -vf crop='ih'",
        "--ppa", "Metadata:"+metadata,
        "-o", outputString,
    )

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        log.Printf("Error executing yt-dlp: %v", err)
        return err
    }

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

    return "app/Music/" + artist + "/" + album + "/" + title + ".%(ext)s"
}