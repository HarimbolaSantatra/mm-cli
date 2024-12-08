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
    fmt.Println("\n" + ruler + "\n")
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
func PrintLineTitle(title, content string) string {
    if strings.Compare(content, "") == 0 {
	// just print the title in green
	return fmt.Sprintf(Green + title + ":", Reset)
    } else {
	// print the title and the content
	return fmt.Sprintf(Green + title + ": " + Reset + Clean(content))
    }
}

// TODO: put this section and the last one into one function
// Print a simple line, with a title in normal text and a content in normal text
func PrintLineSubTitle(title, content string) string {
    if strings.Compare(content, "") == 0 {
	return fmt.Sprintf(BrightGreen + title + ":", Reset)
    } else {
	return fmt.Sprintf(BrightGreen + title + ": " + Reset + Clean(content))
    }
}

// Print in Unordered List format
func PrintUnList(title, content string) string {
    return fmt.Sprintf(Green + title + ": " + Reset + Clean(strings.Replace(content, ",\n", "\n\t- ", -1)))
}

// Print with subsection
// Input:
//  - title string: title of the section
//  - subSectionsTitle string[]: title of each section
//  - subSectionsContent string[]: content of each section
func PrintWithSubSection(title string, subSectionsTitle []string, subSectionsContent []string) string {

    if (len(subSectionsContent) != len(subSectionsTitle)) {
	log.Fatal("Length of the content does not match!")
    }

    result := ""

    result += PrintLineTitle(title, "")

    for i:=0; i < len(subSectionsTitle); i++ {
	result += PrintLineSubTitle(subSectionsTitle[i], subSectionsContent[i])
    }

    result += "eto rahalahy"

    return result

}
