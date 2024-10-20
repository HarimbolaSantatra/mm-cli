package client

import (
    "net/http"
    "io"
    "fmt"
)

func TestIt() {
    endpoint := "https://ifconfig.me"
    resp, _ := http.Get(endpoint)
    defer resp.Body.Close()
    var ip_add []byte
    ip_add, _ = io.ReadAll(resp.Body)
    fmt.Println("Your IP address is", string(ip_add))
}
