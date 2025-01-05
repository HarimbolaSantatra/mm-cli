package client

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"mm/datatype"
	"mm/utils"
)

// Handle error related to Python requirements
func LogPythonError(module string) {
    log.Fatalf("Python module \"%s\" not found! Create a virtual environment and install python packages with 'pip install -r requirements.txt'", module)
}

const scrapingEx = "mm-parsing" // Name of the scraping executable

// Use `./mm-parsing` to get a HTML formatted string and return the JSON format
// Input: a string in HTML format
// Output: the response of the request in JSON format
func ParseString(s string) datatype.ParsedContent {

    basePath, err := os.Getwd()
    if err != nil {
	log.Fatal(err)
    }

    exPath := filepath.Join(basePath, scrapingEx)

    cmd := exec.Command(exPath, s)
    if cmd.Err != nil {
	log.Fatalf("Error in executing %s: %s", scrapingEx, cmd.Err.Error())
    }

    bt, err := cmd.Output()
    if err != nil {
	if err.Error() == "exit status 2" {
	    LogPythonError("bs4")
	}
    }
    return utils.ConvertToParsedContent(bt)
}
