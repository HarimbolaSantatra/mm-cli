package client

import (
    "net/http"
    "io"
    "fmt"
)

const baseUrl = "https://motmalgache.org/bins"
var searchEndpoint = fmt.Sprintf("%s/teny2", baseUrl)

func Search(keyword string) {
    resp, _ := http.Get(searchEndpoint)
    defer resp.Body.Close()
    respBody, _ := io.ReadAll(resp.Body)
    fmt.Println(string(respBody))
}
