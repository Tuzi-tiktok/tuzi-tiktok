package tools

import (
	"io"
	"net/http"
)

type RandomImageResp struct {
	Code   string `json:"code"`
	Imgurl string `json:"imgurl"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

func GetRandomImage() (string, error) {
	resp, err := http.Get("https://www.dmoe.cc/random.php")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Request.URL.String(), nil
}

func GetRandomSentence() (string, error) {
	resp, err := http.Get("https://api.vvhan.com/api/ian")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyText), nil
}
