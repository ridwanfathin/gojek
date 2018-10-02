package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gojek"

	"github.com/labstack/echo"
)

//#account section
func LoginWithEmail(c echo.Context) error {
	data := map[string]string{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("Failed unmarshaling: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := gojek.LoginWithEmail(data["email"])
	fmt.Println(string(res))
	return c.JSON(http.StatusOK, string(res))
}

func LoginWithPhone(c echo.Context) error {
	data := map[string]string{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("Failed unmarshaling: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := gojek.LoginWithPhone(data["phone"])
	fmt.Println(string(res))
	return c.JSON(http.StatusOK, string(res))
}

func CustomerToken(c echo.Context) error {
	data := map[string]string{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("Failed unmarshaling: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := gojek.GenerateCustomerToken(data["otp"], data["token"])
	fmt.Println(string(res))
	return c.JSON(http.StatusOK, string(res))
}

func Logout(c echo.Context) error {

	res := gojek.Logout()
	fmt.Println(string(res))
	return c.JSON(http.StatusOK, string(res))
}

func CustomerInfo(c echo.Context) error {

	res := gojek.GetCustomerInfo()
	fmt.Println(string(res))
	return c.JSON(http.StatusOK, string(res))
}

func EditAccount(c echo.Context) error {
	data := map[string]string{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("Failed unmarshaling: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := gojek.EditAccount(data["phone"], data["email"], data["name"])
	fmt.Println(string(res))
	return c.JSON(http.StatusOK, string(res))
}

func VerifyEditAccount(c echo.Context) error {
	data := map[string]string{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("Failed unmarshaling: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := gojek.EditAccount(data["id"], data["phone"], data["verificationCode"])
	fmt.Println(string(res))
	return c.JSON(http.StatusOK, string(res))
}

//#main section - end
func main() {
	fmt.Println("Welcome to the server")

	e := echo.New()

	//#account section
	e.POST("/loginWithEmail", LoginWithEmail)
	e.POST("/loginWithPhone", LoginWithPhone)
	e.POST("/customer/token", CustomerToken)
	e.DELETE("/customer/logout", Logout)
	e.GET("/customer/info", CustomerInfo)
	e.POST("/customer/edit", EditAccount)
	e.POST("/customer/verifyEdit", VerifyEditAccount)

	e.Start(":8000")
}
