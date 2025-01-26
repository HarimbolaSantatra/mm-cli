package client

import (
    "log"
    "os"
    "os/exec"

    "mm/datatype"
    "mm/utils"
)

// Handle error related to Python requirements
func LogPythonError(module string) {
    log.Fatalf("Python module \"%s\" not found! Create a virtual environment and install python packages with 'pip install -r requirements.txt'", module)
}


const scrapingEx = "mm-parsing"


// Use `./mm-parsing` to get a HTML formatted string and return the JSON format
// Input: string in HTML format
// Output: the response of the request in JSON format
func ParseString(s string) datatype.ParsedContent {

    // Check if parsing script exist
    _, err := exec.LookPath(scrapingEx)
    if err != nil {
	log.Fatalf("%s is not found: %s.\n\nRun 'make install' to install it on PATH.\n\n", scrapingEx, err)
    }

    // Execute command and handle error
    cmd := exec.Command(scrapingEx, s)
    if cmd.Err != nil {
	log.Fatalf("Error in executing %s: %s", scrapingEx, cmd.Err.Error())
    }

    bt, err := cmd.Output()

    if err != nil {
	if err.Error() == "exit status 2" {
	    LogPythonError("bs4")
	}
    }

    // Handle empty output error
    if len(bt) == 0 {
            log.Printf("Input HTML is:\n%s\n\nResult of \"%s\" is empty.\nResult is: '%s'\n", s, scrapingEx, bt)
	// log.Printf("Content of the HTML: %s\n", s)
	os.Exit(1)
    }

    return utils.ConvertToParsedContent(bt)

}
