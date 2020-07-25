package main

import (
	"fmt"

	"github.com/huntergood/tbot/internal/config"
)

var (
	URL string
)

func main() {
	URL = config.GetURL()
	fmt.Println(URL)
}
