package parser

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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
	return string(buf)
}

// GetObject получает объект по указаному селектору
// #TODO:
// .className ищет по классу так же .classsName.className1....
// #id идентификатор #id.className
// tag.className тег с классом
// tag#id тег с идентификатором
// .className tag
// Пока функция принимает готовое регулярное выражение и работает с ним
func GetObject(html, reg string) []string {
	var response = make([]string, 0)
	regular := regexp.MustCompile(reg)
	result := regular.FindAllStringSubmatch(html, -1)
	for _, slice := range result {
		response = append(response, slice[1])
	}
	return response
}
