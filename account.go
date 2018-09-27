package gojek

func LoginWithEmail(email string) []byte {

	var options = map[string]string{
		"method": "POST",
		"body":   `{"email":"` + email + `"}`}

	res := Request(options, "/v3/customers/login_with_email")

	return res
}

func LoginWithPhone(phone string) []byte {
	var options = map[string]string{
		"method": "POST",
		"body":   `{"phone":"` + phone + `"}`}

	res := Request(options, "/v3/customers/login_with_phone")

	return res
}

/**
*  @param otp OTP code from SMS
*  @param loginToken login_token value after calling function loginWithEmail or loginWithPhone
*
*  @return access_token This is customer token. User function gojek.setToken
 */
func GenerateCustomerToken(otp string, loginToken string) []byte {
	var options = map[string]string{
		"method": "POST",
		"body": `{"scope":"gojek:customer:transaction gojek:customer:readonly",` +
			`"grant_type":"password",` +
			`"login_token":"` + loginToken + `",` +
			`"otp":"` + otp + `",` +
			`"client_id":"gojek:cons:android",` +
			`"client_secret":"` + GetClientSecret() + `"}`}

	res := Request(options, "/v3/customers/token")

	return res
}

func Logout() []byte {

	var options = map[string]string{
		"method": "DELETE"}

	res := Request(options, "/v3/auth/token")

	return res
}

func GetCustomerInfo() []byte {

	var options = map[string]string{
		"method": "GET"}

	res := Request(options, "/gojek/v2/customer")

	return res
}

func EditAccount(phone string, email string, name string) []byte {

	var options = map[string]string{
		"method": "POST",
		"body": `{"phone":"` + phone + `",` +
			`{"email":"` + email + `",` +
			`{"name":"` + name + `"}`}

	res := Request(options, "/gojek/v2/customer/edit/v2")

	return res
}

func VerifyEditAccount(id string, phone string, verificationCode string) []byte {

	var options = map[string]string{
		"method": "POST",
		"body": `{"id":"` + id + `",` +
			`"phone":"` + phone + `",` +
			`"verificationCode":"` + verificationCode + `"}`}

	res := Request(options, "/gojek/v2/customer/edit/v2")

	return res
}
