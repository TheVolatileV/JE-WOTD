package main

import (
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
