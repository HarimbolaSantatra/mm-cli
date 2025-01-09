package client

import (
    "os"
    "strings"
    "testing"

    "mm/datatype"
)

func TestParseString(t *testing.T) {

    htmlFile := "../scraping-test/result.html"

    // TODO: this should give error if the file doesn't exist
    b, err := os.ReadFile(htmlFile)
    if err != nil {
	t.Errorf("Error encountered: \"%s\"", err)
    }

    content := ParseString(string(b))

    expected := datatype.ParsedContent{Speech: "Foobar", Vocabulary: "BarBaz"}

    if strings.Compare(content.Speech, expected.Speech) != 0 {
	t.Errorf("Expected %s\nActual: %s", expected.Speech, content.Speech)
    }

    if strings.Compare(content.Vocabulary, expected.Vocabulary) != 0 {
	t.Errorf("Expected %s\nActual: %s", expected.Vocabulary, content.Vocabulary)
    }

}
