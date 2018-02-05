package main

type dict struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []dataStruct `json:"data"`
}

type dataStruct struct {
	Common   bool     `json:"is_common"`
	Tags     []string `json:"tags"`
	Japanese []struct {
		Word    string `json:"word"`
		Reading string `json:"reading"`
	} `json:"japanese"`
	Senses []struct {
		English []string `json:"english_definitions"`
		POS     []string `json:"parts_of_speech"`
		Links   []struct {
			Text string `json:"text"`
			URL  string `json:"url"`
		} `json:"links"`
		Tags         []string `json:"tags"`
		Restrictions []string `json:"restrictions"`
		SeeAlso      []string `json:"see_also"`
		Antonyms     []string `json:"antonyms"`
		Source       []string `json:"source"`
		Info         []string `json:"info"`
	} `json:"senses"`
}

type simpleOutput struct {
	Japanese string   `json:"japanese"`
	English  []string `json:"english"`
	POS      []string `json:"partOfSpeech"`
}
