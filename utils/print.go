package utils

import (
	"fmt"
	"mm/datatype"
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

const mmContributorLink = "https://malagasyword.org/bins/contributors"
const mmHomepage = "https://malagasyword.org/bins/homePage"
const sourceCodeLink = "https://github.com/HarimbolaSantatra/mm-cli"

func PrintRuler() {
    fmt.Println("\n" + ruler)
}

func PrintBanner() {
	PrintRuler()
	fmt.Printf("%s", banner)
	fmt.Println(Red, "Malagasy Word homepage:", Reset, mmHomepage)
	fmt.Println(Red, "Malagasy Word contributors:", Reset, mmContributorLink)
	fmt.Println(Red, "Source code:", Reset, sourceCodeLink)
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
	sb.WriteString(fmt.Sprintf("%s- %s: %s", color, title, Reset))
    } else {
	// print the title and the content
	sb.WriteString(fmt.Sprintf("%s- %s: %s%s", color, title, Reset, Clean(content)))
    }
    return sb.String()
}

// Print an item in a unordered list
func PrintUnListItem(title, content string) string {
    return Green + title + ": " + Reset + Clean(strings.Replace(content, ",\n", "\n\t- ", -1))
}


// Print the final content, which should be the result of the HTML parsing
// Input: an empty interface. This should contain the key `k` and value `v` of the JSON
func PrintResult(parsedContent datatype.ParsedContent, debug bool) {

        fmt.Println(PrintKeyAndValue("Speech Content", parsedContent.Speech, false))
        fmt.Println(PrintKeyAndValue("Malagasy explanation", parsedContent.MlgExplication, false))
        fmt.Println(PrintKeyAndValue("French explanation", parsedContent.FrExplication, false))
        fmt.Println(PrintKeyAndValue("Vocabulary", parsedContent.Vocabulary, false))
        fmt.Println(PrintKeyAndValue("Morphology", parsedContent.Morphology, false))
        fmt.Println(PrintKeyAndValue("Analogs", parsedContent.Analogs, false))

}
