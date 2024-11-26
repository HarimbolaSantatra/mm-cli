package utils

import (
    "fmt"
    "strings"
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

func PrintLineTitle(title string, content string) {
    if strings.Compare(content, "") == 0 {
	fmt.Println(Green + title + ":", Reset)
    } else {
	fmt.Println(Green + title + ": " + Reset + content)
    }
}

// Print in Unordered List format
func PrintUnList(title, content string) {
    PrintLineTitle(title, "")
    fmt.Println(strings.Replace(content, ",\n", "\n\t- ", -1))
}
