package client

import (
    "log"
    "os"
    "os/exec"
    "path/filepath"
)

// Handle error related to Python requirements
func LogPythonError(module string) {
    log.Fatalf("Python module \"%s\" not found! Create a virtual environment and install python packages with 'pip install -r requirements.txt'", module)
}


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
	if err.Error() == "exit status 2" {
	    LogPythonError("bs4")
	}
	if err.Error() == "exit status 3" {
	    LogPythonError("rich")
	}
    }
    return string(bt)
}
