package client

import (
    "net/http"
    "io"
    "fmt"
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

func Search(keyword string) string {
    resp, _ := http.Get(searchEndpoint)
    defer resp.Body.Close()
    respBody, _ := io.ReadAll(resp.Body)
    return string(respBody)
}
