package main

import (
	"fmt"

	"news.com/events/helper"
)

const immoScout24BaseUrl = "https://www.immobilienscout24.de/Suche/de/"


func main() {
	x := helper.ImmoScoutRegion[0]
	fmt.Println(x)
}