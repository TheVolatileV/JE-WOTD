package main

import (
	"encoding/json"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello World!"

	m, err := json.Marshal(str)
	if err != nil {
		panic(err)
	}
	w.Write(m)
}
