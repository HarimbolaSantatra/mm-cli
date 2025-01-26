package utils

import (
	"fmt"
	"strings"
	"testing"
)

// Function used by test function to print debugging information
func DebugPrint(actual, expected string) string {
	return fmt.Sprintf("Got: \"%s\"\nExpecting: \"%s\"", actual, expected)
}

func TestClean(t *testing.T) {
	const x = "Ito dia texte[3.45] fanao[v.45] test[ okon. ]ma80nde tsara v?]"
	y := Clean(x)
	expected := "Ito dia texte fanao testmande tsara v?]"
	if strings.Compare(y, expected) != 0 {
		t.Errorf(DebugPrint(y, expected))
	}
}

func TestPrintKeyAndValue(t *testing.T) {
	const title = "Soratra maintso"
	const content = "soratra fotsy"
	expected := fmt.Sprintf("%s- %s: %s%s", Green, title, Reset, content)
	x := PrintKeyAndValue(title, content, false)
	if strings.Compare(x, expected) != 0 {
		t.Errorf("\nGot: \"%s\"\nExpecting: \"%s\"", x, expected)
	}
}
