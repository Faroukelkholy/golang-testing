package locations_provider

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"golang-testing/src/api/domain/locations"
	"golang-testing/src/api/utils/errors"
	"net/http"
)

const (
	url           = "https://api.mercadolibre.com"
	urlGetCountry = url + "/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))

	if response == nil || response.Response == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient error when getting country %s", countryId),
			Error:   "",
		}
	}

	if response.StatusCode > 299 {
		var apiErr errors.ApiError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
				Error:   "",
			}
		}
		return nil, &apiErr
	}

	var result locations.Country

	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
			Error:   "",
		}
	}

	return &result, nil
}
