package test

import (
	"log"
	"strconv"
	"testing"
)

func TestSplice(t *testing.T) {
	var strs = ([]string)(nil)
	//strs := []string{
	//	"1", "2", "3",
	//}
	ints := make([]int, len(strs))
	for i := range strs {
		t, err := strconv.Atoi(strs[i])
		if err != nil {
			panic(err)
		}
		ints[i] = t
	}
	log.Println(ints)
}
