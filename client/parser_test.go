package client

import (
    "strings"
    "testing"

    "mm/datatype"
)

func TestParseString(t *testing.T) {

    htmlFile := "../result.html"

    content := ParseString(string(htmlFile), true)

    expected := datatype.ParsedContent{Speech: "2noun", Vocabulary: "5Economy: food", MlgExplication: "3[1.1]Trano fanaovana nahandro:Tsy miala ao an-dakozia foana io saka io", FrExplication: "4[1.3][lakoZY] cuisine"}

    if strings.Compare(content.Speech, expected.Speech) != 0 {
	t.Errorf("Expected %s\nActual: %s", expected.Speech, content.Speech)
    }

    if strings.Compare(content.MlgExplication, expected.MlgExplication) != 0 {
	t.Errorf("Expected %s\nActual: %s", expected.MlgExplication, content.MlgExplication)
    }

    if strings.Compare(content.FrExplication, expected.FrExplication) != 0 {
	t.Errorf("Expected %s\nActual: %s", expected.FrExplication, content.FrExplication)
    }

    if strings.Compare(content.Vocabulary, expected.Vocabulary) != 0 {
	t.Errorf("Expected %s\nActual: %s", expected.Vocabulary, content.Vocabulary)
    }

}
