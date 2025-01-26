package client

import (
    "strings"
    "testing"

    "mm/datatype"
)

func TestParseString(t *testing.T) {

    htmlFile := "../result.html"

    content := ParseString(string(htmlFile))

    expected := datatype.ParsedContent{Speech: "2noun", Vocabulary: "5Economy: food"}

    if strings.Compare(content.Speech, expected.Speech) != 0 {
	t.Errorf("Expected %s\nActual: %s", expected.Speech, content.Speech)
    }

    if strings.Compare(content.Vocabulary, expected.Vocabulary) != 0 {
	t.Errorf("Expected %s\nActual: %s", expected.Vocabulary, content.Vocabulary)
    }

}
