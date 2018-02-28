package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	jaWord := obj.Data[0].Japanese[0].Word
	reading := obj.Data[0].Japanese[0].Reading

	var engWords []string
	if len(obj.Data[0].Senses[0].English) >= 4 {
		engWords = obj.Data[0].Senses[0].English[:4]
	} else {
		engWords = obj.Data[0].Senses[0].English
	}

	pos := obj.Data[0].Senses[0].POS
	japPos := getPartOfSpeech(pos)
	out := simpleOutput{
		jaWord,
		reading,
		engWords,
		japPos,
	}
	json.NewEncoder(w).Encode(out)
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

func getPartOfSpeech(engPos []string) []string {
	soln := make([]string, 0)
	japPosMap := map[string]string{"adverbial noun": "名詞", "adverb": "副詞", "verb": "動詞",
		"noun": "名詞", "adjective": "形容詞", "particle": "助詞"}
	partOfSpeech, _ := regexp.Compile("(?i)(adverbial noun|adjective|noun|verb|adverb|particle)")

	for _, element := range engPos {
		if partOfSpeech.MatchString(element) {
			japPos := japPosMap[strings.ToLower(partOfSpeech.FindString(element))]
			if !inListAlready(japPos, soln) {
				soln = append(soln, japPos)
			}
		}
	}
	return soln
}

func inListAlready(japPos string, soln []string) bool {
	for _, elem := range soln {
		if elem == japPos {
			return true
		}
	}
	return false
}
