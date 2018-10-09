package gojek

var token = ""
var uniqueId = "9774d56d682e549c"
var appVersion = "2.28.2"
var clientSecret = "83415d06-ec4e-11e6-a41b-6c40088ab51e"
var location = "-6.180495106.824992"

func SetToken(Token string) {
	token = Token
}
func GetToken() string {
	return token
}
func SetUniqueId(UniqueId string) {
	uniqueId = UniqueId
}
func GetUniqueId() string {
	return uniqueId
}
func SetAppVersion(AppVersion string) {
	appVersion = AppVersion
}
func GetAppVersion() string {
	return appVersion
}
func GetClientSecret() string {
	return clientSecret
}
func SetLocation(Location string) {
	location = Location
}
func GetLocation() string {
	return location
}
