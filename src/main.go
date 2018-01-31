package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/justinas/alice"
	"github.com/rs/cors"
)

func main() {
	startup()
	waitForShutdown()
}

func startup() {
	tcpPort := 80
	startServer(tcpPort)
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

func getRandomWord() string {
	resp, err := http.Get("https://nlp.fi.muni.cz/projekty/random_word/run.cgi?language_selection=en&word_selection=verbs&model_selection=use&length_selection=&probability_selection=true")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	word := string(body)
	return word
}
