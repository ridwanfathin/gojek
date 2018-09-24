# gojek
Un-official Go-jek API Wrapper written in GO. Rewrited from https://github.com/mychaelgo/gojek

Documentation
=============

## Getting Started

You need `go` installed and set GOPATH to your project. Once that is done, run the
command:
```shell
$ go get -u github.com/ridwanfathin/gojek
```

## Configuration
By default the module set the `location`, `uniqueId` and `appVersion`. This value used to every request to the Go-Jek API. You can set manually if you need.

### Set unique id
```go
gojek.setUniqueId('YOUR_UNIQUE_ID');
```

### Set app version
```go
gojek.setAppVersion('YOUR_APP_VERSION');
```

### Set token
By default the token is not set by this module. You can set token after you call a login API
```go
gojek.setToken('YOUR_TOKEN');
```


## Account
### Login
Go-Jek support 2 method for login (Email or Phone number login)

```go
gojek.loginWithEmail('your@email.com', function(error ,response, body){
	console.log(body);
});
```

```go
gojek.loginWithPhone('+628123456789', function(error ,response, body){
	console.log(body);
});
```

After request that API, the registered phone number will receive an OTP. 
You must save your `login_token` to be used in next step :

```go
gojek.generateCustomerToken('1234', 'login_token', function(error ,response, body){
	console.log(body);
});
```

Save `access_token`, then call :
```go
gojek.setToken('access_token');
```

### Get customer info
```go
gojek.getCustomerInfo(function (err, res, body) {
    console.log(body);
});
```

### Edit account
- Param 1: Phone
- Param 2: Email
- Param 3: Name
```go
gojek.editAccount('+628123456789','email@domain.com','NAME', function (err, res, body) {
    console.log(body);
});
```


### Logout

## Go-Pay
### Get Go-Pay info
```go
gojek.getGoPayDetail(function (err, res, body) {
    console.log(body);
});
```
### Get Go-Pay transaction history 
- Param 1: Page number (start from 1)
- Param 2: Limit per page
```go
gojek.getGoPayHistory(1, 30, function (err, res, body) {
    console.log(body);
});
```
### Get Go-Pay id by phone
```go
gojek.getGoPayQrId('+628123456789', function (err, res, body) {
    console.log(body);
});
```
### Transfer Go-Pay money
```go
gojek.transferGoPay('QR_ID', 10000, 'YOUR_DESCRIPTION', function (err, res, body) {
    console.log(body);
});
```

## Go-Mart
### Get nearest Go-Mart
- Param 1: latitude,longitude
```go
gojek.getNearestGoMart('-6.180495,106.824992', function (err, res, body) {
    console.log(body);
});
```

## Go-Food
### Get Go-Food Home
```go
gojek.setToken('ACCESS_TOKEN');
gojek.setLocation('-6.180495,106.824992');
gojek.getGoFoodHome(gojek.getLocation(), function (err, res, body) {
    console.log(body);
});
```
### Get nearest Go-Food
- Param 1: latitude,longitude
- Param 2: Page (start from 0)
- Param 3: Limit
```go
gojek.getNearestGoFood(gojek.getLocation(), 0, 10, function (err, res, body) {
    console.log(body);
});
```

### Get restaurant
- Param 1: Restaurant UUID
```go
gojek.getRestaurant('UUID', function (err, res, body) {
    console.log(body);
});
```

### Get restaurants by category
- Param 1: Category code (Can be seen on get go-food home)
- Param 2: Page number (start from 0)
- Param 3: Limit per page
```go
gojek.setToken('ACCESS_TOKEN');
gojek.setLocation('-6.180495,106.824992');
gojek.getRestaurantsByCategory('HEALTHY_FOOD', '0', '32', function (err, res, body) {
    console.log(body);
});
```

## Booking
### Get active booking
```go 
gojek.getActiveBooking(function (err, res, body) {
    console.log(body);
});
```

### Get booking history
```go
gojek.getBookingHistory(function (err, res, body) {
    console.log(body);
});
```

### Get booking by order no
```go
gojek.getBookingByOrderNo('123456', function (err, res, body) {
    console.log(body);
});
```

### Cancel booking

## Go-Points

### Get Go-Points
```go
gojek.getGoPoints(function (err, res, body) {
    console.log(body);
});
```

### Next Go-Points
```go
gojek.nextPointsToken(function (err, res, body) {
    console.log(body);
});
```

### Redeem Go-Points
```go
gojek.redeemGoPointsToken('POINTS_TOKEN_ID',function (err, res, body) {
    console.log(body);
});
```