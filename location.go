package gojek

func ReverseGeocode(latLong string) []byte {

	var options = map[string]string{
		"method": "GET",
		"qs":     `{"location":"` + latLong + `"}`}

	res := Request(options, "/gojek/poi/reverse-geocode/")
	return res
}

func GetGoRideNearby(latLong string) []byte {

	var options = map[string]string{
		"method": "GET",
		"qs":     `{"location":"` + latLong + `"}`}

	res := Request(options, "/gojek/service_type/1/drivers/nearby")
	return res
}

func GetGoCarNearby(latLong string) []byte {

	var options = map[string]string{
		"method": "GET",
		"qs":     `{"location":"` + latLong + `"}`}

	res := Request(options, "/gojek/service_type/13/drivers/nearby")
	return res
}

func GetGoSendNearby(latLong string) []byte {

	var options = map[string]string{
		"method": "GET",
		"qs":     `{"location":"` + latLong + `"}`}

	res := Request(options, "/gojek/service_type/14/drivers/nearby")
	return res
}
