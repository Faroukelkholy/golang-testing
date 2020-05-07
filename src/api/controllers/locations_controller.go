package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-testing/src/api/services"
	"log"
	"net/http"
)

func GetCountry(c *gin.Context) {
	country, err := services.ILocationService.GetCountry(c.Param("country_id"))
	log.Println("country",country)
	log.Println("err",err)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, country)
}
