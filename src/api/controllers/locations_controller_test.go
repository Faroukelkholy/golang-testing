package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang-testing/src/api/domain/locations"
	"golang-testing/src/api/services"
	"golang-testing/src/api/utils/errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	url           = "https://api.mercadolibre.com"
	urlGetCountry = url + "/countries/%s"
)

var (
	getCountryMockFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

type locationServiceMock struct{}

func (*locationServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryMockFunc(countryId)
}

func TestMain(m *testing.M) {
	//restclient.StartMockups()
	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
	//countryId := "AR"
	// mocking restclient in locations_provider
	//restclient.FlushMockups()
	//restclient.AddMockup(restclient.Mock{
	//	Url:        fmt.Sprintf(urlGetCountry, countryId),
	//	HttpMethod: http.MethodGet,
	//	Response: &http.Response{
	//		StatusCode: http.StatusNotFound,
	//		Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Country not found","error":"not_found","status":404,"cause":[]}`)),
	//	},
	//	Err: nil,
	//})
	getCountryMockFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    "Country not found",
			Error:      "",
		}
	}
	services.ILocationService = &locationServiceMock{}
	//creating a context required by the controller and pass a response struct to record response data,in order to be used with testing
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	//c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiErr errors.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.StatusCode)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}


func TestGetCountry(t *testing.T) {
	getCountryMockFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{Id: "AR", Name: "Argentina"} , nil
	}
	services.ILocationService = &locationServiceMock{}

	response :=httptest.NewRecorder()

	c,_ := gin.CreateTestContext(response)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	var location locations.Country

	err := json.Unmarshal(response.Body.Bytes(),&location)
	assert.Nil(t,err)
	assert.NotNil(t,location)

	assert.EqualValues(t,http.StatusOK,response.Code)
	assert.EqualValues(t,"AR",location.Id)
	assert.EqualValues(t,"Argentina",location.Name)

}