package gojek

func GetGoPayDetail() []byte {
	var options = map[string]string{
		"method": "GET"}

	res := Request(options, "/wallet/profile/detailed")
	return res
}

// page start from 1
func GetGoPayHistory(page string, limit string) []byte {

	var options = map[string]string{
		"method": "GET",
		"qs": `{"page":"` + page + `",` +
			`"client_secret":"` + limit + `"}`}

	res := Request(options, "/wallet/history")

	return res
}

func GetGoPayQrId(phoneNumber string) []byte {

	var options = map[string]string{
		"method": "GET",
		"qs":     `{"phone_number":"` + phoneNumber + `"}`}

	res := Request(options, "/wallet/qr-code")

	return res
}

func TransferGoPay(qrId string, amount string, description string, pinCode string) []byte {

	var options = map[string]string{
		"method": "POST",
		"qs": `{"qr_id":"` + qrId + `",` +
			`"amount":"` + amount + `",` +
			`"description":"` + description + `"}`,
		"headers": `{"pin":"` + pinCode + `"}`}

	res := Request(options, "/v2/fund/transfer")

	return res
}
