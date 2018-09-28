package gojek

func GetNearestGoMart(latLong string) []byte {

	var options = map[string]string{
		"method": "GET",
		"qs":     `{"location":"` + latLong + `"}`}

	res := Request(options, "/gojek/mart-category/listMartNearest")
	return res
}
