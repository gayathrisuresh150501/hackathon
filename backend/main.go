package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	lr "github.com/LoginRadius/go-sdk"
	"github.com/savsgio/atreugo/v11"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}

func main() {
	client()

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

	// Define a simple POST route
	server.Path("POST", "/greet", func(ctx *atreugo.RequestCtx) error {
		name := string(ctx.PostArgs().Peek("name"))
		message := fmt.Sprintf("Hello, %s!", name)
		return ctx.TextResponse(message)
	})

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
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
