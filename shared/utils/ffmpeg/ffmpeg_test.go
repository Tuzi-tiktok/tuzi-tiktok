package ffmpeg

import (
	"io"
	"os"
	"testing"
)

func TestGetSnapShots(t *testing.T) {
	shots, err := GetSnapShots("http://127.0.0.1:9000/images-storage/TMP_STREAM_PATH/64a020112c3b4a883f6dd290.mp4")
	if err != nil {
		t.Fatal(err)
		return
	}
	file, err := os.OpenFile("snap.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatal(err)
		return
	}
	_, err = io.Copy(file, shots)
	if err != nil {
		t.Fatal(err)
		return
	}
}
