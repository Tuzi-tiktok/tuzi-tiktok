package ffmpeg

import (
	"bytes"
	ff "github.com/u2takey/ffmpeg-go"
	"io"
)

// GetSnapShots ffmpeg -ss 00:03:00 -i vURL -c:v png -f image2pipe -vframes 1 pipe:1
func GetSnapShots(vURL string) (io.Reader, error) {
	input := ff.Input(vURL, ff.KwArgs{
		"ss": "00:00:30",
	})
	buffer := bytes.NewBuffer(nil)
	args := ff.KwArgs{
		"vframes": 1,
		"f":       "image2pipe",
		"c:v":     "png",
	}
	err := input.Output("pipe:1", args).WithOutput(buffer).Run()
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
