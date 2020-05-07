package locations_provider

import (
	"encoding/json"
	"fmt"
	"golang-testing/src/api/clients/restclient"
	"golang-testing/src/api/domain/locations"
	"golang-testing/src/api/utils/errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	url           = "https://api.mercadolibre.com"
	urlGetCountry = url + "/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	log.Println("location_provider Getcountry countryId", countryId)
	response, err := restclient.Get(fmt.Sprintf(urlGetCountry, countryId), http.Header{}, nil)
	//log.Println("locationProvider.GetCountry response  ", response)
	//log.Println("locationProvider.GetCountry err  ", err)
	if err != nil {
		println("err in provider.getCountry", err)
	}
	if response == nil || response.Body == nil {
		return nil, &errors.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("invalid restclient error when getting country %s", countryId),
			Error:      "",
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &errors.ApiError{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var apiErr errors.ApiError
		if err := json.Unmarshal(bytes, &apiErr); err != nil {
			return nil, &errors.ApiError{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("invalid error interface when getting country %s", countryId),
				Error:      "",
			}
		}
		apiErr.StatusCode = response.StatusCode
		return nil, &apiErr
	}

	var result locations.Country

	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, &errors.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
			Error:      "",
		}
	}

	return &result, nil
}
