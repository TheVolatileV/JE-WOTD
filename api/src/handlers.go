package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getWord() dict {
	word := getRandomWord()
	fmt.Println(word)
	if word == "is" || word == "has" || word == "have" {
		word = getRandomWord()
	}
	yandexWord := translateWord(word)
	resp, err := http.Get(fmt.Sprintf("http://jisho.org/api/v1/search/words?keyword=%s", url.PathEscape(yandexWord)))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	obj := dict{}
	if err = json.Unmarshal(body, &obj); err != nil {
		fmt.Println(err)
	}
	return obj
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := getWord()

	out := subsetJisho(data)
	json.NewEncoder(w).Encode(out)
}

func subsetJisho(data dict) simpleOutput {
	if len(data.Data) == 0 {
		newData := getWord()
		fmt.Println("was empty")
		return subsetJisho(newData)
	}
	jaWord := data.Data[0].Japanese[0].Word
	reading := data.Data[0].Japanese[0].Reading

	var engWords []string
	if len(data.Data[0].Senses[0].English) >= 4 {
		engWords = data.Data[0].Senses[0].English[:4]
	} else {
		engWords = data.Data[0].Senses[0].English
	}

	pos := data.Data[0].Senses[0].POS
	return simpleOutput{
		jaWord,
		reading,
		engWords,
		pos,
	}
}

func translateWord(word string) string {
	//use yandex to translate the word
	type yandex struct {
		Text string `xml:"text"`
	}
	resp, err := http.Get(fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr/translate?key=%s&text=%s&lang=ja", apiKey, strings.ToLower(word)))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var jaWord yandex
	err = xml.Unmarshal(body, &jaWord)
	if err != nil {
		panic(err)
	}
	fmt.Println(jaWord.Text)
	return jaWord.Text
}
