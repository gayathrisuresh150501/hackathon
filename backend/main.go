package main

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	lr "github.com/LoginRadius/go-sdk"
	"github.com/savsgio/atreugo/v11"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}

func main() {
	// Server configuration
	config := atreugo.Config{
		Addr: "0.0.0.0:8080", // Listening on port 8080
	}

	// Create a new Atreugo instance
	server := atreugo.New(config)

	// Define a simple GET route
	server.Path("GET", "/hello", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Hello, Atreugo!")
	})

	server.Path("POST", "/login", func(ctx *atreugo.RequestCtx) error {
		body := ctx.PostBody()

		var user models.UserEmailPassword

		json.Unmarshal(body, &user)

		loginResp := loginWithEmail(user.Email, user.Password)

		return ctx.TextResponse(loginResp)
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func loginWithEmail(email, password string) string {
	loginurl := "https://devapi.lrinternal.com/identity/v2/auth/login"

	// build payload
	login := models.UserEmailPassword{
		Email:    email,
		Password: password,
	}
	payloadByte, _ := json.Marshal(login)

	payload := bytes.NewReader(payloadByte)

	apikey := os.Getenv("APIKEY")

	// set querry params
	data := url.Values{}
	data.Set("apikey", apikey)

	// build request with querry params
	req, _ := http.NewRequest("POST", loginurl+"?"+data.Encode(), payload)

	fmt.Println(string(payloadByte))

	req.Header.Add("content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func client() {
	apiKey := "26c87dc1-3b2d-41f7-9688-2405ab7c3c88"
	apiSecret := "e95a3fa8-7ee9-4ca9-9e95-1f003e0cdbd0"
	email := "gayathri.suresh@loginradius.com"

	urlStr := "https://devapi.lrinternal.com/identity/v2/manage/account/identities"
	data := url.Values{}
	data.Set("apikey", apiKey)
	data.Set("apisecret", apiSecret)
	data.Set("email", email)

	req, err := http.NewRequest("GET", urlStr+"?"+data.Encode(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(res)
	fmt.Println(string(body))
}
