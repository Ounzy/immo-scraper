package main

import (
	"fmt"

	"strings"

	"github.com/gocolly/colly" // import Colly
)

func main() {

	var name string // variable for what we scrape (maybe also type struct)

	var nameList []string // if there is more than 1 of this class: List of the variable/struct from above

	c := colly.NewCollector() // initialise Colly

	c.OnHTML("[data-container-type=\"geohierarchieeinstieg\"] li", func(e *colly.HTMLElement) { // choose the element we want to scrape

		name = e.ChildText("a") // choose the attributes we want to scrape

		name = strings.ToLower(name)
		name = strings.Trim(name, " ")
		r := strings.NewReplacer("ö", "oe", "ü", "ue", "ä", "ae", "()", "", " ", "-", "/", "-")
		name = r.Replace(name)

		nameList = append(nameList, name) // append data to list

	})

	c.Visit("https://www.immobilienscout24.de/wohnen/hamburg,hamburg.html") // what website do we want to visit/scrape

	fmt.Println(nameList)

	for _, s := range nameList {
		fmt.Println("{\"hamburg\", \"hamburg,\", \"" + s + "\"},")
	}
}
