package utils

import (
    "fmt"
    "strings"
    "log"
)

// Print a JSON content, which should be the result of the HTML parsing
// Input: key `k` and value `v` of the JSON
func PrintResult(fmtResp interface{}) {
    m := fmtResp.(map[string]interface{})

    for k, v := range m {

	// Remove eventual whitespaces 
	trimedK := strings.TrimSpace(k)

	switch vv := v.(type) {
	    case string:

		// Split into multiple line for section 'Analogues'
		if (strings.Compare(trimedK, "Analogues") == 0) {
		    fmt.Println(PrintUnList(trimedK, vv))
		    break
		}

		// Default print mode for all section except these mentioned above
		// Do not print 'Mots composés'
		if (strings.Compare(trimedK, "Mots composés") != 0) {
		    fmt.Println(PrintLineTitle(trimedK, vv))
		    break
		}

		// Section `Morphologie` contains a sub section
		if (strings.Compare(trimedK, "Morphologie") == 0) {
		    // Get the JSON beneath it
		    fmt.Printf("%s\n", PrintWithSubSection(trimedK, vv))
		    break
		}

	    default:
		log.Fatalf("%s dia type tsy mbola hitako hatr@izay niainana!", k)

	    }

    }

}
