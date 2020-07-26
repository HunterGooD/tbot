package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetHTML получает всю страницу
func GetHTML(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "text/html")
	//  пока задается напрямую
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
	return string(buf)
}

// GetObject получает объект по указаному селектору
// #TODO:
// .className ищет по классу так же .classsName.className1....
// #id идентификатор #id.className
// tag.className тег с классом
// tag#id тег с идентификатором
// .className tag
func GetObject(selector string) []string {
	return []string{}
}

// map[string]string {"class": "nameClass", "tag": ""}
func getSelection(selector string) map[string]string {
	return make(map[string]string)
}
