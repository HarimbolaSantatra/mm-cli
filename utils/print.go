package utils

import (
    "fmt"
    "strings"
    "regexp"
)

var Reset = "\033[0m" 
var Red = "\033[31m" 
var Green = "\033[32m" 
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
    fmt.Println(Red, "Malagasy Word homepage:", Reset , mmHomepage)
    fmt.Println(Red, "Malagasy Word contributors:", Reset, mmContributorLink)
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

// Print a simple line
func PrintLineTitle(title string, content string) {
    if strings.Compare(content, "") == 0 {
	// just print the title in green
	fmt.Println(Green + title + ":", Reset)
    } else {
	// print the title and the content
	fmt.Println(Green + title + ": " + Reset + Clean(content))
    }
}

// Print in Unordered List format
func PrintUnList(title, content string) {
    PrintLineTitle(title, "")
    fmt.Println(Clean(strings.Replace(content, ",\n", "\n\t- ", -1)))
}
