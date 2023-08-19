package ffmpeg

import (
	"bytes"
	ff "github.com/u2takey/ffmpeg-go"
	"io"
	"net/http"
	"sync"
)

var (
	single = sync.Once{}
	c      *http.Client
)

func init() {

}
func InitializeClient() {
	single.Do(func() {
		c = &http.Client{}
	})
}

// GetSnapShots ffmpeg -ss 00:00:03 -i vURL -c:v png -f image2pipe -vframes 1 pipe:1
func GetSnapShots(vURL string) (io.Reader, error) {
	input := ff.Input(vURL, ff.KwArgs{
		"ss": "00:00:03",
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
	if buffer.Len() == 0 {
		return RandomCover()
	}

	return buffer, nil
}
func RandomCover() (io.Reader, error) {
	if c == nil {
		InitializeClient()
	}
	resp, err := c.Get("https://www.dmoe.cc/random.php")
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(make([]byte, resp.ContentLength))
	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
