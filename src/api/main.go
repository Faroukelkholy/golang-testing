package main

import (
	"fmt"
	"golang-testing/src/api/providers/locations_provider"
)

func main() {
	//log.Println("main func")
	//test := flag.String("test", "", "test var")
	//flag.Parse()
	//if *test != "" {
	//	log.Println("test :", *test)
	//}
	country, err := locations_provider.GetCountry("AR")
	fmt.Println(country)
	fmt.Println(err)
}
