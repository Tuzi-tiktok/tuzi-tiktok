package ffmpeg

import (
	"io"
	"log"
	"net/http"
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
func TestHttpClient(t *testing.T) {
	client := http.Client{}
	response, err := client.Get("https://www.dmoe.cc/random.php")
	if err != nil {
		t.Fatal(err)
		return
	}
	code := response.StatusCode
	log.Println(code)
	file, err := os.OpenFile("a.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	io.Copy(file, response.Body)
}
