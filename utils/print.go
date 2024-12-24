package utils

import (
    "fmt"
    "log"
    "regexp"
    "strings"
)

var Reset = "\033[0m" 
var Red = "\033[31m" 
var Green = "\033[32m" 
var BrightGreen = "\033[92m" 
var White = "\033[97m"

const banner = `
 #    # #    #        ####  #      # 
 ##  ## ##  ##       #    # #      # 
 # ## # # ## # ##### #      #      # 
 #    # #    #       #      #      # 
 #    # #    #       #    # #      # 
 #    # #    #        ####  ###### # 

`
const ruler = "===================================="

const mmContributorLink = "https://motmalgache.org/bins/contributors"
const mmHomepage = "https://motmalgache.org/bins/homePage"
const sourceCodeLink = "https://github.com/HarimbolaSantatra/mm-cli"

func GetVersion() string {
    version := "0.1b"
    return version
}

func PrintRuler() {
    fmt.Println("\n" + ruler)
}

func PrintBanner() {
    PrintRuler()
    fmt.Printf("%s", banner)
    fmt.Println(Red, "Monde Malgache homepage:", Reset , mmHomepage)
    fmt.Println(Red, "Monde Malgache contributors:", Reset, mmContributorLink)
    fmt.Println(Red, "Source code:", Reset, sourceCodeLink)
    fmt.Println(Red, "Version:", Reset, GetVersion())
    PrintRuler()
}

// Remove source reference (format `[...]` from a string) and numbers
// Input: string
// Return a new string
func Clean(txt string) string {
    const rx = "\\[.+?\\]|\\d+?" // regex for [..] and numbers
    re := regexp.MustCompile(rx)
    return re.ReplaceAllString(txt, "")
}

// Print a simple line with a title in bold and a content in normal text
func PrintKeyAndValue(title, content string, isSubTitle bool) string {

    var sb strings.Builder

    var color string
    if isSubTitle {
	color = BrightGreen
    } else { 
	color = Green
    }

    if strings.Compare(content, "") == 0 {
	// just print the title in color and ignore the content
	sb.WriteString(fmt.Sprintf("%s%s: %s", color, title, Reset))
    } else {
	// print the title and the content
	sb.WriteString(fmt.Sprintf("\t%s- %s: %s%s", color, title, Reset, Clean(content)))
    }
    return sb.String()
}

// Print an item in a unordered list
func PrintUnListItem(title, content string) string {
    return Green + title + ": " + Reset + Clean(strings.Replace(content, ",\n", "\n\t- ", -1))
}


// Print the final content, which should be the result of the HTML parsing
// Input: an empty interface. This should contain the key `k` and value `v` of the JSON
func PrintResult(inp interface{}, debug bool) {

    // Assert and convert `inp` to a map
    m := inp.(map[string]interface{})

    for k, v := range m {

	// Remove eventual whitespaces 
	trimedK := strings.TrimSpace(k)

	switch vv := v.(type) {
	    case string:

		// Split into multiple line for section 'Analogues'
		if (strings.Compare(trimedK, "Analogues") == 0) {
		    fmt.Println(PrintUnListItem(trimedK, vv))
		    break
		}

		// Default print mode for all section except these mentioned above
		// Do not print 'Mots composés'
		if (strings.Compare(trimedK, "Mots composés") != 0) {
		    fmt.Println(PrintKeyAndValue(trimedK, vv, false))
		    break
		}

	    case interface{}:
		// Section `Morphologie` contains a sub section and should use this
		keys, values := GetKeysValues(vv, debug)

		if debug {
		    log.Fatalf("Value of keys and value in the PrintResult:\n%s\n%s\n", keys, values)
		}

		// fmt.Printf("%s\n", PrintSubSection(trimedK, keys, values))
		break

	    default:
		log.Fatalf("%s dia type tsy mbola hitako hatr@izay niainana!", k)

	    }

    }

}
