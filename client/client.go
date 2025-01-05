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

const baseUrl = "https://malagasyword.org/bins"
var homepageEnd = fmt.Sprintf("%s/homePage", baseUrl)
var searchEndpoint = fmt.Sprintf("%s/teny2", baseUrl)

func Health() int {
    resp, err := http.Get(homepageEnd)
    if err != nil {
	log.Fatalf("Error in making request: %s", err)
    }
    var resCode int;
    if resp.StatusCode > 299 {
	resCode = 3
	log.Printf("Status code was %d", resp.StatusCode)
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

// Make a request to malagasyword.org
// Input: string, a search keyword
// Return: string, the html content of the result
func Search(keyword string) string {

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

    // Close
    defer resp.Body.Close()

    return string(body)
}
