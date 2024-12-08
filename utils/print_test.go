package utils

import (
	"strings"
	"testing"
)

func TestPrintWithSubSection(t *testing.T) {

    title := "User"
    sTitle := []string { "first name", "last name"}
    sContent := []string { "Richard", "Rapport"}

    // Build string for the expected value
    var sb strings.Builder
    sb.WriteString(Green + title + ": " + Reset + "\n")
    sb.WriteString("\t" + BrightGreen + "- first name: " + Reset + "Richard\n")
    sb.WriteString("\t" + BrightGreen + "- last name: " + Reset + "Rapport\n")
    expected := sb.String()

    got := PrintWithSubSection(title, sTitle, sContent)
    if strings.Compare(got, expected) != 0 {
	t.Errorf(DebugPrint(got, expected))
    }
}
