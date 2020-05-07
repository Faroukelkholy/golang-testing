package locations_provider

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-testing/src/api/clients/restclient"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	countryId = "AR"
)


func TestGetCountryrestclientError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        fmt.Sprintf(urlGetCountry, countryId),
		HttpMethod: http.MethodGet,
		Response: &http.Response{
		},
	})
	country, err := GetCountry(countryId)
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid restclient error when getting country AR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        fmt.Sprintf(urlGetCountry, countryId),
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Country not found","error":"not_found","status":404,"cause":[]}`)),
		},
	})
	//restclient.AddMockup(&restclient.Mock{
	//	Url:          fmt.Sprintf(urlGetCountry, countryId),
	//	HttpMethod:   http.MethodGet,
	//	RespHTTPCode: http.StatusNotFound,
	//	RespBody:     `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	//})

	country, err := GetCountry(countryId)

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        fmt.Sprintf(urlGetCountry, countryId),
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":Country not found,"error":"not_found","status":404,"cause":[]}`)),
		},
	})
	//restclient.AddMockup(&restclient.Mock{
	//	Url:          fmt.Sprintf(urlGetCountry, countryId),
	//	HttpMethod:   http.MethodGet,
	//	RespHTTPCode: http.StatusNotFound,
	//	RespBody:     `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	//})

	country, err := GetCountry(countryId)

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid error interface when getting country AR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        fmt.Sprintf(urlGetCountry, countryId),
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123,"name":"Argentina","time_zone":"GMT-03:00"}`)),
		},
	})

	//restclient.AddMockup(&restclient.Mock{
	//	Url:          fmt.Sprintf(urlGetCountry, countryId),
	//	HttpMethod:   http.MethodGet,
	//	RespHTTPCode: http.StatusNotFound,
	//	RespBody:     `{"id":123,"name":"Argentina","time_zone":"GMT-03:00"}`,
	//})

	country, err := GetCountry(countryId)

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error when trying to unmarshal country data for AR", err.Message)
}

func TestGetCountry(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        fmt.Sprintf(urlGetCountry, countryId),
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":"AR","name":"Argentina","locale":"es_AR","currency_id":"ARS","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-38.416096,"longitude":-63.616673}},"states":[{"id":"AR-B","name":"Buenos Aires"},{"id":"AR-C","name":"Capital Federal"},{"id":"AR-K","name":"Catamarca"},{"id":"AR-H","name":"Chaco"},{"id":"AR-U","name":"Chubut"},{"id":"AR-W","name":"Corrientes"},{"id":"AR-X","name":"Córdoba"},{"id":"AR-E","name":"Entre Ríos"},{"id":"AR-P","name":"Formosa"},{"id":"AR-Y","name":"Jujuy"},{"id":"AR-L","name":"La Pampa"},{"id":"AR-F","name":"La Rioja"},{"id":"AR-M","name":"Mendoza"},{"id":"AR-N","name":"Misiones"},{"id":"AR-Q","name":"Neuquén"},{"id":"AR-R","name":"Río Negro"},{"id":"AR-A","name":"Salta"},{"id":"AR-J","name":"San Juan"},{"id":"AR-D","name":"San Luis"},{"id":"AR-Z","name":"Santa Cruz"},{"id":"AR-S","name":"Santa Fe"},{"id":"AR-G","name":"Santiago del Estero"},{"id":"AR-V","name":"Tierra del Fuego"},{"id":"AR-T","name":"Tucumán"}]}`)),
		},
	})
	//restclient.AddMockup(&restclient.Mock{
	//	Url:          fmt.Sprintf(urlGetCountry, countryId),
	//	HttpMethod:   http.MethodGet,
	//	RespHTTPCode: http.StatusNotFound,
	//	RespBody:     `{"id":"AR","name":"Argentina","locale":"es_AR","currency_id":"ARS","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-38.416096,"longitude":-63.616673}},"states":[{"id":"AR-B","name":"Buenos Aires"},{"id":"AR-C","name":"Capital Federal"},{"id":"AR-K","name":"Catamarca"},{"id":"AR-H","name":"Chaco"},{"id":"AR-U","name":"Chubut"},{"id":"AR-W","name":"Corrientes"},{"id":"AR-X","name":"Córdoba"},{"id":"AR-E","name":"Entre Ríos"},{"id":"AR-P","name":"Formosa"},{"id":"AR-Y","name":"Jujuy"},{"id":"AR-L","name":"La Pampa"},{"id":"AR-F","name":"La Rioja"},{"id":"AR-M","name":"Mendoza"},{"id":"AR-N","name":"Misiones"},{"id":"AR-Q","name":"Neuquén"},{"id":"AR-R","name":"Río Negro"},{"id":"AR-A","name":"Salta"},{"id":"AR-J","name":"San Juan"},{"id":"AR-D","name":"San Luis"},{"id":"AR-Z","name":"Santa Cruz"},{"id":"AR-S","name":"Santa Fe"},{"id":"AR-G","name":"Santiago del Estero"},{"id":"AR-V","name":"Tierra del Fuego"},{"id":"AR-T","name":"Tucumán"}]}`,
	//})
	country, err := GetCountry(countryId)

	assert.NotNil(t, country)
	assert.Nil(t, err)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 24, len(country.States))
}
