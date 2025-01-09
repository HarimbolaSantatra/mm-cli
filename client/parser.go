package client

import (
    "errors"
    "fmt"
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


const scrapingEx = "mm-parsing" // Name of the scraping executable


// Use `./mm-parsing` to get a HTML formatted string and return the JSON format
// Input: a string in HTML format
// Output: the response of the request in JSON format
func ParseString(s string) datatype.ParsedContent {

    // TODO: edit this relative path
    exPath := fmt.Sprintf("../%s", scrapingEx)

    // Check if the executable exist
    if _, err := os.Stat(exPath); errors.Is(err, os.ErrNotExist) {
	log.Fatalf("Executable %s does not exist", exPath)
    }

    // Execute command and handle error
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

    // Handle empty output error
    if len(bt) == 0 {
	log.Printf(fmt.Sprintf("Result of \"%s\" is empty", exPath))
	// log.Printf("Content of the HTML: %s\n", s)
	os.Exit(1)
    }

    return utils.ConvertToParsedContent(bt)

}
