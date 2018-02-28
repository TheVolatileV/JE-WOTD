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

const apiKey = "trnsl.1.1.20180207T225403Z.5c84170a42a76523.0dbfe4e01c06ed2e14a1ed63227253eb3f60e35d"

var words = wordsDict{}

func main() {
	startup()
	waitForShutdown()
}

func startup() {
	tcpPort, _ := strconv.Atoi(os.Getenv("PORT"))
	if tcpPort == 0 {
		tcpPort = 80
	}
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
	err = json.Unmarshal(body, &words)
	if err != nil {
		panic(err)
	}
}

func getRandomWord() string {
	return words.Data[rand.Intn(len(words.Data))].Word
}
