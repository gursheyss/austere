package ytdlp

import (
	"fmt"
	"os/exec"
)

func Download() error {
	url := "https://www.youtube.com/watch?v=YG3EhWlBaoI"

	cmd := exec.Command("yt-dlp", url, "-x", "--audio-format", "mp3", "--add-metadata",
		"--parse-metadata", "artist:%(artist||uploader)s",
		"--parse-metadata", "title:%(title)s",
		"--parse-metadata", "album:%(album)s",
		"--replace-in-metadata", "album:NA=%(title)s",
		"--ppa", "ThumbnailsConvertor+ffmpeg_o:-c:v png -vf crop='ih'",
		"--embed-thumbnail",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing yt-dlp:", err)
	}
	fmt.Println(string(output))
	return nil
}
