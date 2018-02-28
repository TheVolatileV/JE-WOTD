package main

import "testing"

func TestSubsetJisho(t *testing.T) {
	populateWords()
	resp := getWord()
	subset := subsetJisho(resp)
	if subset.Japanese != resp.Data[0].Japanese[0].Word {
		t.Errorf("Did not subset properly")
	}
}

func TestGetWord(t *testing.T) {
	populateWords()
	resp := getWord()
	if len(resp.Data) == 0 {
		t.Errorf("Did not successfully get a word")
	}
}
