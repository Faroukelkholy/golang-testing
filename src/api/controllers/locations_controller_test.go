package controllers

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}



func TestGetCountryNotFound(t *testing.T){
	//mock rest request

	c,_ := gin.CreateTestContext()
}