package gojek

func loginWithEmail(email string) {

	Request("POST", "/v3/customers/login_with_email")
}
