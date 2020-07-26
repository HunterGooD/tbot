package parser

// GetHTML получает всю страницу
func GetHTML(url string) string {
	return ""
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
