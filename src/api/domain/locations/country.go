package locations

type Country struct {
	Id             string
	Name           string
	TimeZone       string         `json:"time_zone"`
	GeoInformation GeoInformation `json:"geo_information"`
	States         []State
}

type GeoInformation struct {
	Location GeoLocation
}

type GeoLocation struct {
	Latitude  float64
	Longitude float64
}

type State struct {
	Id   string
	Name string
}
