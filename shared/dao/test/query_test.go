package test

import (
	"log"
	"testing"
	"tuzi-tiktok/dao/query"
)

func TestQuery(t *testing.T) {
	users, err := query.User.Find()
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}
	for _, user := range users {
		log.Printf("%v", user)
	}
}
