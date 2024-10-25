package client

import (
    "net/http"
    "io"
    "fmt"
    "log"
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

// Data for search
type Data struct {
    keyword string
}

func (data Data) Read(p []byte) (n int, err error) {
    return len(p), err
}

func Search(keyword string) string {
    data := Data{ keyword }
    resp, err := http.Post(searchEndpoint, "text/html", data)
    defer resp.Body.Close()
    if err != nil {
	log.Fatal("Error in posting to motmalgache!")
    }
    respBody, _ := io.ReadAll(resp.Body)
    return string(respBody)
}
