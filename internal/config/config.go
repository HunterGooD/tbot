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
	DBname string `json:"dbname"`
}

// NewSpecial ..
func NewSpecial() *Special {
	var s = &Special{}
	b, err := ioutil.ReadFile("configs/special.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(b, s); err != nil {
		log.Fatal(err)
	}

	return s
}

// URL ..
func (s *Special) URL() string {
	return s.URLAPI + "bot" + s.Token
}
