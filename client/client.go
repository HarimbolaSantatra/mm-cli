package client

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
    "strings"
    "time"
)

const baseUrl = "https://motmalgache.org/bins"
var homepageEnd = fmt.Sprintf("%s/homePage", baseUrl)
var searchEndpoint = fmt.Sprintf("%s/teny2", baseUrl)

func Health() int {
    resp, err := http.Get(homepageEnd)
    var resCode int;
    if resp.StatusCode > 299 || err != nil {
	resCode = 1
    } else {
	resCode = 0
    }
    defer resp.Body.Close()
    return resCode
}

// request Data
type Data struct {
    requestData string
}

func (data Data) Read(p []byte) (n int, err error) {
    return len(p), err
}

func Search(keyword string) {

    // Create Client
    client := &http.Client{
	Timeout: time.Second * 15,
    }

    // Set request data
    data := url.Values{}
    data.Add("w", keyword)
    encodedData := data.Encode()


    // Make request
    resp, err := client.Post(searchEndpoint,
	"application/x-www-form-urlencoded",
	strings.NewReader(encodedData))
    if err != nil {
	log.Fatalf("Request Error: %s", err.Error())
    }

    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))

    // Close
    defer resp.Body.Close()
}
