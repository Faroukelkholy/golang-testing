package services

import (
	"golang-testing/src/api/domain/locations"
	"golang-testing/src/api/providers/locations_provider"
	"golang-testing/src/api/utils/errors"
	"log"
)

var (
	ILocationService ILocationServiceInterface
)

type ILocationServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}
type locationService struct{}

func init() {
	ILocationService = &locationService{}
}
func (*locationService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	log.Println("location_service Getcountry countryId", countryId)
	location, apiErr := locations_provider.GetCountry(countryId)
	if apiErr != nil {
		log.Println("location_service apiErr", apiErr)
		return nil, apiErr
	}
	log.Println("location_service location", location)
	return location, nil
}
