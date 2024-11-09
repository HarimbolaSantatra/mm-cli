package parser

import (
    "log"
    "os/exec"
)

// Get a HTML formatted string and return the JSON format
// This function uses the python script `./scraping` to do that.
// Input: a string in HTML format
// Output: the response of the request in JSON format
func ParseString(s string) string {
    cmd := exec.Command("./scraping", s)
    bt, err := cmd.Output()
    if err != nil {
	log.Fatalf("Command execution error: ", err)
    }
    return string(bt)
}
