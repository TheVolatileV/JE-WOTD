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

func TestGetPartOfSpeech(t *testing.T) {
	ex := []string{"adverbial noun", "adverb", "verb", "noun", "adjective", "particle"}
	answer := []string{"名詞", "副詞", "動詞", "形容詞", "助詞"}
	ex = getPartOfSpeech(ex)
	if len(ex) != len(answer) {
		t.Errorf("Did not produce same number of parts of speech")
	}
	for i := 0; i < len(ex); i++ {
		if ex[i] != answer[i] {
			t.Errorf("Did not successfully translate part of speech")
		}
	}
}
