package main

import (
	"strings"
	"testing"
)

func TestSubsetJisho(t *testing.T) {
	populateWords()
	resp := getWord()
	subset := subsetJisho(resp)
	if subset.Japanese != resp.Data[0].Japanese[0].Word &&
		strings.Join(subset.English, ",") != strings.Join(resp.Data[0].Senses[0].English, ",") &&
		subset.Reading != resp.Data[0].Japanese[0].Reading &&
		strings.Join(subset.POS, ",") != strings.Join(resp.Data[0].Senses[0].POS, ",") {
		t.Errorf("Did not subset properly")
	}
}

func TestGetWord(t *testing.T) {
	populateWords()
	resp := getWord()
	if len(resp.Data) == 0 && resp.Meta.Status == 200 {
		t.Errorf("Did not successfully get a word")
	}
}
