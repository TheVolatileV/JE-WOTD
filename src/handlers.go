package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://jisho.org/api/v1/search/words?keyword=house")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	obj := dict{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj.Data[0].Japanese[0].Word)
	w.Write(body)
}
