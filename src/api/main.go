package main

import (
	"fmt"
	"golang-testing/src/api/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("AR")
	fmt.Println(country)
	fmt.Println(err)
}
