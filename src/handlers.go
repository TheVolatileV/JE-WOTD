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
	wordWithNL := getRandomWord()
	word := strings.Trim(wordWithNL, "\n")
	resp, err := http.Get(fmt.Sprintf("http://jisho.org/api/v1/search/words?keyword=%s", word))
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
	engWord := obj.Data[0].Senses[0].English[0]
	pos := obj.Data[0].Senses[0].POS
	out := simpleOutput{
		jaWord,
		engWord,
		pos,
	}
	json.NewEncoder(w).Encode(out)
}
