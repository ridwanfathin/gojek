package gojek

func GetGoPoints() []byte {
	var options = map[string]string{
		"method": "GET"}

	res := Request(options, "/gopoints/v1/wallet/points-balance")
	return res
}

func NextPointsToken() []byte {

	var options = map[string]string{
		"method": "POST",
		"body":   "{}"}

	res := Request(options, "/gopoints/v1/next-points-token")
	return res
}

func RedeemGoPointsToken(goPointsToken string) []byte {

	var options = map[string]string{
		"method": "POST",
		"body":   `{"points_token_id":"` + goPointsToken + `"}`}

	res := Request(options, "/gopoints/v1/redeem-points-token")
	return res
}
