package gojek

/**
    * @param latLong Required latitude,longitude
    */
func GetGoFoodHome(latLong string) []byte {
    var options = map[string]string{
        "method": "GET",
        "qs":`{"location":"` + latLong + `"}`}

	res := Request(options, "/gofood/consumer/v2/home")
	return res
}

// page start from 0
func GetNearestGoFood(latLong string, page string, limit string) []byte {

    var options = map[string]string{
		"method": "GET",
		"qs": `{"location":"` + latLong + `",` +
			`"page":"` + page + `",` +
			`"client_secret":"` + limit + `"}`}

	res := Request(options, "/gojek/merchant/find")

	return res
}

func GetRestaurant(restaurantId string) []byte {

    var options = map[string]string{
		"method": "GET"}

	res := Requestoptions, "/gofood/consumer/v2/restaurants/" + restaurantId)

	return res
}

func GetRestaurantsByCategory(category string, page string, limit string) []byte {

    var options = map[string]string{
		"method": "GET",
		"qs": `{"category":"` + latLong + `",` +
			`"page":"` + page + `",` +
			`"client_secret":"` + limit + `"}`}

	res := Request(options, "/gofood/consumer/v2/restaurants")

	return res
}