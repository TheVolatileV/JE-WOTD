package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/justinas/alice"
	"github.com/rs/cors"
)

var Words = wordsDict{}

func main() {
	startup()
	waitForShutdown()
}

func startup() {
	tcpPort := 80
	startServer(tcpPort)
	populateWords()
}

func startServer(tcpPort int) {

	log.Print("Starting server")
	router := NewRouter()
	corsInstance := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		Debug:            false,
	})
	pipeline := alice.New(corsInstance.Handler).Then(router)

	log.Print("Server listening on TCP port " + strconv.Itoa(tcpPort))
	go func() {
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(tcpPort), pipeline))
	}()
}

func waitForShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}

func populateWords() {
	url := "https://randomwordgenerator.com/json/words.json"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &Words)
	if err != nil {
		panic(err)
	}
}

// can do verbs, adjectives, and nouns
// verbs, adjecs, nouns respectively
func getRandomWord() string {
	//url := fmt.Sprintf("https://nlp.fi.muni.cz/projekty/random_word/run.cgi?language_selection=en&word_selection=%s&model_selection=use&length_selection=&probability_selection=true",
	//	"verbs")
	return Words.Data[rand.Intn(len(Words.Data))].Word
}
