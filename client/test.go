package client

import (
    "testing"
    "mm/utils"
)

func TestSearch(t *testing.T) {
    const htmlFile = "scraping-test/result.html"
    const keyword = "lakozia"
    result := Search(keyword)
    res := utils.fvCompare(htmlFile, result)
    if !res {
	t.Error("error testing the Search() function!")
    }
}
