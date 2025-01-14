package utils

import (
	"fmt"
	"strings"
)

// Print a JSON content, which should be the result of the HTML parsing
// Input: key `k` and value `v` of the JSON
func PrintResult(fmtResp interface{}) {
	m := fmtResp.(map[string]interface{})
	for k, v := range m {
		trimedK := strings.TrimSpace(k) // Remove enventual whitespaces
		switch vv := v.(type) {
		case string:

			// 'Morphologie' contains multiple strings
			if strings.Compare(trimedK, "Morphologie") == 0 {
				PrintUnList(trimedK, vv)
			} else {

				// Remove section 'Mots composés' because it's too long!
				// Maybe handle this in the future if I'm not lazy!
				if strings.Compare(trimedK, "Mots composés") != 0 {

					// Split into multiple line for section 'Analogues'
					if strings.Compare(trimedK, "Analogues") == 0 {
						PrintUnList(trimedK, vv)
					} else {
						// Default print mode for all section except these mentioned above
						PrintLineTitle(trimedK, vv)
					}

				}
			}

		default:
			fmt.Println(k, " dia type tsy mbola hitako hatr@izay niainana!")
		}
	}
}
