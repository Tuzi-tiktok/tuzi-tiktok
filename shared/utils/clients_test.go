package utils

import "testing"

func TestNewAuth(t *testing.T) {
	_, err := NewAuth()
	if err != nil {
		panic(err)
	}
}
