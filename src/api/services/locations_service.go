package services

import (
	"golang-testing/src/api/domain/locations"
	"golang-testing/src/api/providers/locations_provider"
	"golang-testing/src/api/utils/errors"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
