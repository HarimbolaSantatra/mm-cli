package client

import (
    "net/http"
    "io"
    "fmt"
)

func TestIt() {
    endpoint := "https://httpbin.org/get"
    resp, _ := http.Get(endpoint)
    defer resp.Body.Close()
    var ip_add []byte
    ip_add, _ = io.ReadAll(resp.Body)
    fmt.Println(string(ip_add))
}
