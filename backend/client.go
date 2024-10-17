package main

// func client() {
// 	apiKey := "26c87dc1-3b2d-41f7-9688-2405ab7c3c88"
// 	apiSecret := "a567ac93-eb2a-48cb-b16f-be3a95d935ab"
// 	email := "gayathri.suresh@loginradius.com"

// 	urlStr := "https://config.lrcontent.com/ciam/appinfo"
// 	data := url.Values{}
// 	data.Set("apikey", apiKey)
// 	data.Set("apisecret", apiSecret)
// 	data.Set("email", email)

// 	req, err := http.NewRequest("GET", urlStr+"?"+data.Encode(), nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Println("Error making request:", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	fmt.Println(res)
// 	fmt.Println(string(body))
// }
