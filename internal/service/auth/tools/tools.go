package tools

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type RandomImageResp struct {
	Code   string `json:"code"`
	Imgurl string `json:"imgurl"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

func GetRandomImage() (string, error) {
	resp, err := http.Get("https://www.dmoe.cc/random.php?return=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data RandomImageResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(data.Imgurl, "\\", ""), nil
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
