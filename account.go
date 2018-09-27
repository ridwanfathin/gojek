package gojek

func LoginWithEmail(email string) []byte {
	var jsonStr = []byte(`{"email":"` + email + `"}`)

	res := Request("POST", jsonStr, "/v3/customers/login_with_email")

	return res
}

func LoginWithPhone(phone string) []byte {
	var jsonStr = []byte(`{"phone":"` + phone + `"}`)

	res := Request("POST", jsonStr, "/v3/customers/login_with_phone")

	return res
}

/**
*  @param otp OTP code from SMS
*  @param loginToken login_token value after calling function loginWithEmail or loginWithPhone
*
*  @return access_token This is customer token. User function gojek.setToken
 */
func GenerateCustomerToken(otp string, loginToken string) []byte {
	var jsonStr = []byte(
		`{"scope":"gojek:customer:transaction gojek:customer:readonly",` +
			`"grant_type":"password",` +
			`"login_token":"` + loginToken + `",` +
			`"otp":"` + otp + `",` +
			`"client_id":"gojek:cons:android",` +
			`"client_secret":"` + GetClientSecret() + `"}`)

	res := Request("POST", jsonStr, "/v3/customers/token")

	return res
}

func Logout() []byte {

	res := Request("DELETE", nil, "/v3/auth/token")

	return res
}

func GetCustomerInfo() []byte {

	res := Request("GET", nil, "/gojek/v2/customer")

	return res
}

func EditAccount(phone string, email string, name string) []byte {

	var jsonStr = []byte(
		`{"phone":"` + phone + `",` +
			`{"email":"` + email + `",` +
			`{"name":"` + name + `"}`)

	res := Request("POST", jsonStr, "/gojek/v2/customer/edit/v2")

	return res
}

func VerifyEditAccount(id string, phone string, verificationCode string) []byte {

	var jsonStr = []byte(
		`{"id":"` + id + `",` +
			`"phone":"` + phone + `",` +
			`"verificationCode":"` + verificationCode + `"}`)

	res := Request("POST", jsonStr, "/gojek/v2/customer/edit/v2")

	return res
}
