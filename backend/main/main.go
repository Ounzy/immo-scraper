package main

import (
	"fmt"
	"strconv"

	"news.com/events/entities"
	"news.com/events/helper"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

//const immoScout24BaseUrl = "https://www.immobilienscout24.de/Suche/de/"

func main() {
	x := helper.ImmoScoutRegion[12000]
	fmt.Println(x)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))

	router.GET("/immoScoutRegions", getImmoScoutRegions)
	router.POST("/search", postSearchParameters)

	router.Run("localhost:8080")
}

//
// get the whole immoscoutregionlist

func getImmoScoutRegions(c *gin.Context) {
	var regions []entities.Region

	for _, region := range helper.ImmoScoutRegion {
		regions = append(regions, entities.Region{
			Bundesland: region[0],
			Kreis:      region[1],
			Stadt:      region[2],
		})
	}

	c.JSON(http.StatusOK, regions)
}

func postSearchParameters(c *gin.Context) {
	var searchParamters entities.SearchParamter
	if err := c.BindJSON(&searchParamters); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messages": "Invalid input!"})
	}

	fmt.Println(searchParamters)

	var urlToFetch = "https://www.immobilienscout24.de/Suche/de/" + searchParamters.Where[1] + "/" + searchParamters.Where[2] + "/" + searchParamters.Where[0] + "/"

	if searchParamters.What != "" && searchParamters.Kind != "" {
		urlToFetch = urlToFetch + searchParamters.What + "-" + searchParamters.Kind + "?"
	} else {
		fmt.Println("funnkt nicht")
	}

	if searchParamters.RoomNumber != 0 {
		urlToFetch = urlToFetch + "numberofrooms=" + strconv.FormatFloat(searchParamters.RoomNumber, 'f', 1, 64)
	}

	if searchParamters.PriceEnd != 0 && searchParamters.PriceStart != 0 {
		urlToFetch = urlToFetch + "-&price=" + strconv.FormatFloat(searchParamters.PriceStart, 'f', 2, 64) + "-" + strconv.FormatFloat(searchParamters.PriceEnd, 'f', 2, 64)
	} else if searchParamters.PriceStart != 0 {
		urlToFetch = urlToFetch + "-&price=" + strconv.FormatFloat(searchParamters.PriceStart, 'f', 2, 64)
	} else if searchParamters.PriceEnd != 0 {
		urlToFetch = urlToFetch + "-&price=-" + strconv.FormatFloat(searchParamters.PriceEnd, 'f', 2, 64)
	}

	if searchParamters.Squaremeters != 0 {
		urlToFetch = urlToFetch + "&livingspace=" + strconv.Itoa(searchParamters.Squaremeters)
	}

	// https://www.immobilienscout24.de/Suche/de/baden-wuerttemberg/rems-murr-kreis/weinstadt/wohnung-mieten?numberofrooms=2.0-&price=-300000.0&livingspace=10.0

	fmt.Println(urlToFetch)

}
