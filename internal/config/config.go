package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type special struct {
	URLAPI string `json:"urlApi"`
	Token  string `json:"token"`
}

// GetURL Возвращает url для общения с API
func GetURL() string {
	var spec = special{}
	b, err := ioutil.ReadFile("configs/special.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(b, &spec); err != nil {
		log.Fatal(err)
	}
	return spec.concatURL()
}

func (s *special) concatURL() string {
	return s.URLAPI + "bot" + s.Token
}
