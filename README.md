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
gojek.SetUniqueId('YOUR_UNIQUE_ID')
```

### Set app version
```go
gojek.SetAppVersion('YOUR_APP_VERSION')
```

### Set token
By default the token is not set by this module. You can set token after you call a login API
```go
gojek.SetToken('YOUR_TOKEN')
```


## Account
### Login
Go-Jek support 2 method for login (Email or Phone number login)

```go
res := gojek.LoginWithEmail('your@email.com') //[]byte
```

```go
res := gojek.LoginWithPhone('+628123456789') //[]byte
```

After request that API, the registered phone number will receive an OTP. 
You must save your `login_token` to be used in next step :

```go
res := gojek.GenerateCustomerToken('1234', 'login_token') //[]byte
```

Save `access_token`, then call :
```go
gojek.SetToken('access_token')
```

### Get customer info
```go
res := gojek.GetCustomerInfo()
```

### Edit account
- Param 1: Phone
- Param 2: Email
- Param 3: Name
```go
res := gojek.EditAccount('+628123456789','email@domain.com','NAME')
```


### Logout

## Go-Pay
### Get Go-Pay info
```go
res := gojek.GetGoPayDetail()
```
### Get Go-Pay transaction history 
- Param 1: Page number (start from 1)
- Param 2: Limit per page
```go
gojek.GetGoPayHistory(1, 30)
```
### Get Go-Pay id by phone
```go
res := gojek.getGoPayQrId('+628123456789')
```
### Transfer Go-Pay money
```go
res := gojek.TransferGoPay('QR_ID', 10000, 'YOUR_DESCRIPTION')
```

## Go-Mart
### Get nearest Go-Mart
- Param 1: latitude,longitude
```go
res := gojek.GetNearestGoMart('-6.180495,106.824992')
```

## Go-Food
### Get Go-Food Home
```go
gojek.SetToken('ACCESS_TOKEN')
gojek.SetLocation('-6.180495,106.824992')
res := gojek.GetGoFoodHome(gojek.getLocation())
```
### Get nearest Go-Food
- Param 1: latitude,longitude
- Param 2: Page (start from 0)
- Param 3: Limit
```go
res := gojek.GetNearestGoFood(gojek.getLocation(), 0, 10)
```

### Get restaurant
- Param 1: Restaurant UUID
```go
res := gojek.GetRestaurant('UUID')
```

### Get restaurants by category
- Param 1: Category code (Can be seen on get go-food home)
- Param 2: Page number (start from 0)
- Param 3: Limit per page
```go
gojek.SetToken('ACCESS_TOKEN')
gojek.SetLocation('-6.180495,106.824992')
res := gojek.GetRestaurantsByCategory('HEALTHY_FOOD', '0', '32')
```

## Booking
### Get active booking
```go 
res := gojek.GetActiveBooking()
```

### Get booking history
```go
res := gojek.GetBookingHistory()
```

### Get booking by order no
```go
res := gojek.GetBookingByOrderNo('123456')
```

### Cancel booking

## Go-Points

### Get Go-Points
```go
res := gojek.GetGoPoints()
```

### Next Go-Points
```go
res := gojek.NextPointsToken()
```

### Redeem Go-Points
```go
res := gojek.RedeemGoPointsToken('POINTS_TOKEN_ID')
```