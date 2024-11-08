package utils

import (
    "os"
    "bufio"
    "log"
)

// Compare if a file have the same content as a variable
// Arguments: filename and the content of the variable
func fvCompare(filename string, compVar string) bool {

    f, err := os.Open(filename)
    if err != nil {
	log.Fatal(err)
    }

    fscan := bufio.NewScanner(f)

    for fscan.Scan() {
	if fscan.Text() == compVar  {
	    return true
	}
    }

    return false

}
