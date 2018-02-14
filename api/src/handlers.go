package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	word := getRandomWord()
	fmt.Println(word)
	if word == "is" || word == "has" || word == "have" {
		word = getRandomWord()
	}
	resp, err := http.Get(fmt.Sprintf("http://jisho.org/api/v1/search/words?keyword=%s", strings.ToLower(word)))
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	obj := dict{}
	if err = json.Unmarshal(body, &obj); err != nil {
		fmt.Println(err)
	}
	jaWord := obj.Data[0].Japanese[0].Word
	reading := obj.Data[0].Japanese[0].Reading

	var engWords []string
	if len(obj.Data[0].Senses[0].English) >= 4 {
		engWords = obj.Data[0].Senses[0].English[:4]
	} else {
		engWords = obj.Data[0].Senses[0].English
	}

	pos := obj.Data[0].Senses[0].POS
	out := simpleOutput{
		jaWord,
		reading,
		engWords,
		pos,
	}
	json.NewEncoder(w).Encode(out)
}

func translateWord(word string) {
	//use yandex to translate the word
}
