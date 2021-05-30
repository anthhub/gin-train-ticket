package searchcitylist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/anthhub/gin-train-ticket/citylist"

	"github.com/gin-gonic/gin"
	"github.com/mozillazg/go-pinyin"
)

// SearchCityResult ....
type SearchCityResult struct {
	Result    []SearchResult `json:"result"`
	SearchKey string         `json:"searchKey"`
}

// SearchResult ....
type SearchResult struct {
	Key     string `json:"key"`
	Display string `json:"display"`
}

// SearchCityList ...
func SearchCityList(c *gin.Context) {
	key := strings.Trim(c.DefaultQuery("key", ""), " ")

	fmt.Println(key, "")

	file, _ := ioutil.ReadFile("cities.json")

	data := citylist.CityResult{}

	_ = json.Unmarshal([]byte(file), &data)

	hotCities := data.HotCities

	var resultHot []SearchResult
	for _, value := range hotCities {
		resultHot = append(resultHot, SearchResult{Key: value.Name, Display: value.Name})
	}

	a := pinyin.NewArgs()

	chars := strings.Split(key, "")
	var charsChineseList []string

	for _, value := range chars {
		charArr := pinyin.LazyPinyin(value, a)
		char := value
		if len(charArr) > 0 {
			char = charArr[0]
		}
		charsChineseList = append(charsChineseList, char)
	}

	charsChineseStr := strings.Join(charsChineseList, "")
	firstChineseChar := strings.ToUpper(strings.Split(charsChineseStr, "")[0])

	fmt.Println("charsChineseList \n ", charsChineseList)

	fmt.Println("firstChineseChar \n ", firstChineseChar)

	// 选出名称首字母相同的
	var resultMatchedCities citylist.CityList
	for _, value := range data.CityList {
		if firstChineseChar == value.Title {
			resultMatchedCities = value
			break
		}
	}

	fmt.Println("resultMatchedCities \n ", resultMatchedCities)

	// 选出名字包含的
	var resultMatched []SearchResult
	for _, value := range resultMatchedCities.Citys {

		chars := strings.Split(value.Name, "")
		var charsChineseList []string

		for _, value := range chars {
			charArr := pinyin.LazyPinyin(value, a)
			char := value
			if len(charArr) > 0 {
				char = charArr[0]
			}
			charsChineseList = append(charsChineseList, char)
		}
		nameChineseStr := strings.Join(charsChineseList, "")

		if strings.Contains(charsChineseStr, nameChineseStr) || strings.Contains(nameChineseStr, charsChineseStr) {

			resultMatched = append(resultMatched, SearchResult{Key: value.Name, Display: value.Name})
		}
	}

	// fmt.Printf("resultMatched=%v resultHot=%v \n ", resultMatched, resultHot)

	c.JSON(http.StatusOK, append(resultMatched, resultHot...))
}
