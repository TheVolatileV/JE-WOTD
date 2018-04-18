package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var lastExecuted time.Time
var refreshMin, _ = time.ParseDuration("24h")
var currentWord simpleOutput

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if time.Now().After(lastExecuted.Add(refreshMin)) {
		data := getWord()
		currentWord = subsetJisho(data)
		lastExecuted = time.Now()
	}

	json.NewEncoder(w).Encode(currentWord)
}

func forceNewWordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := getWord()
	currentWord = subsetJisho(data)
	lastExecuted = time.Now()

	json.NewEncoder(w).Encode(currentWord)
}

func insertEmailHandler(w http.ResponseWriter, r *http.Request) {
	var email emailAddr
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	if err := json.Unmarshal(body, &email); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	err = db.insertEmail(email.Email, tableName)
}

func getEmailListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list, err := db.scanTable(tableName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(list)
}

func getWord() dict {
	word := getRandomWord()
	fmt.Println(word)
	if word == "is" || word == "has" || word == "have" {
		word = getRandomWord()
	}
	yandexWord := translateWord(strings.ToLower(word))
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

func subsetJisho(data dict) simpleOutput {
	if len(data.Data) == 0 {
		newData := getWord()
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
	japPos := getPartOfSpeech(pos)
	return simpleOutput{
		jaWord,
		reading,
		engWords,
		japPos,
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
