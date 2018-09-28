package gojek

func GetBookingHistory(userId string) []byte {
	var options = map[string]string{
		"method": "GET",
	}

	res := Request(options, "/gojek/v2/booking/history/"+userId)
	return res
}

func GetActiveBooking() []byte {
	var options = map[string]string{
		"method": "GET",
	}

	res := Request(options, "/v1/customers/active_bookings")
	return res
}

func GetBookingByOrderNo(orderNo string) []byte {
	var options = map[string]string{
		"method": "GET",
	}

	res := Request(options, "/gojek/v2/booking/findByOrderNo/"+orderNo)
	return res
}

func Calculate() []byte {
	var options = map[string]string{
		"method": "POST",
	}

	res := Request(options, "/gojek/v2/calculate/gopay/")
	return res
}
