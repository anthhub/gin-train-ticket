package main

import (
	"net/http"

	"github.com/anthhub/gin-train-ticket/citylist"
	"github.com/anthhub/gin-train-ticket/searchcitylist"

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
	// Listen and Server in 0.0.0.0:8000
	r.Run(":8000")
}
