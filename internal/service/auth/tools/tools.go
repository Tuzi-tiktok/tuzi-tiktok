package tools

import (
	"io"
	"net/http"
)

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
