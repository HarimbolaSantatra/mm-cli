package client

import (
    "log"
    "os"
    "os/exec"
    "path/filepath"
)

// Get a HTML formatted string and return the JSON format
// This function uses the python script `./scraping` to do that.
// Input: a string in HTML format
// Output: the response of the request in JSON format
func ParseString(s string) string {

    basePath, err := os.Getwd()
    if err != nil {
	log.Fatal(err)
    }

    exPath := filepath.Join(basePath, "scraping")

    cmd := exec.Command(exPath, s)
    bt, err := cmd.Output()
    if err != nil {
	log.Fatalf("Command execution error: %s", err)
    }
    return string(bt)
}
