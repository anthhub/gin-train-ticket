package main

import (
	"net/http"

	"basic/citylist"
	"basic/searchcitylist"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/api/cities", citylist.GetCityList)

	r.GET("/api/search", searchcitylist.SearchCityList)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
