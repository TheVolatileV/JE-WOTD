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
		English      []string `json:"english_definitions"`
		POS          []string `json:"parts_of_speech"`
		Links        []string `json:"links"`
		Tags         []string `json:"tags"`
		Restrictions []string `json:"restrictions"`
		SeeAlso      []string `json:"see_also"`
		Antonyms     []string `json:"antonyms"`
		Source       []string `json:"source"`
		Info         []string `json:"info"`
	} `json:"senses"`
	Attribution struct {
		JMdict   bool `json:"jmdict"`
		JMNedict bool `json:"jmnedict"`
		DBpedia  bool `json:"dbpedia"`
	} `json:"attribution"`
}
