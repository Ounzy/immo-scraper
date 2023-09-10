package main

import (
	"fmt"

	"news.com/events/entities"
	"news.com/events/helper"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

const immoScout24BaseUrl = "https://www.immobilienscout24.de/Suche/de/"

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
	// Suchergebnisse scrapen
}
