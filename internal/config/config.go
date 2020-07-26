package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Special config
type Special struct {
	URLAPI string `json:"urlApi"`
	Token  string `json:"token"`
	DbNAme string `json:"dbname"`
}

// GetURL Возвращает url для общения с API
func (s *Special) GetURL() string {
	b, err := ioutil.ReadFile("configs/special.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(b, s); err != nil {
		log.Fatal(err)
	}
	return s.concatURL()
}

func (s *Special) concatURL() string {
	return s.URLAPI + "bot" + s.Token
}
