package citylist

import (
	"fmt"
	"net/http"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// CityResult ...
type CityResult struct {
	HotCities []HotCities `json:"hotCities"`
	CityList  []CityList  `json:"cityList"`
	Version   int         `json:"version"`
}

// HotCities ...
type HotCities struct {
	Name string `json:"name"`
}

// CityList ...
type CityList struct {
	Title string `json:"title"`
	Citys []City `json:"citys"`
}

// City ...
type City struct {
	Name string `json:"name"`
}

// GetCityList ...
func GetCityList(c *gin.Context) {
	file, err1 := ioutil.ReadFile("./static/cities.json")

	if err1 != nil {
		fmt.Printf("ReadFile err: %v; ", err1)
	}

	data := CityResult{}

	err := json.Unmarshal([]byte(file), &data)

	if err != nil {
		fmt.Printf("Unmarshal err: %v; ", err)
	}

	// fmt.Printf("data: %v; ", data)

	c.JSON(http.StatusOK, data)
}
