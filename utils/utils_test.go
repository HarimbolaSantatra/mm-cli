package utils

import (
    "strings"
    "testing"

    "mm/datatype"
)

func TestingCompare(t *testing.T, str1, str2 string) {
    if strings.Compare(str1, str2) != 0 {
	t.Errorf("%s and %s => Not the same!", str1, str2)
    } else
    {
	t.Logf("%s and %s => OK!", str1, str2)
    }
}

func TestConvertToParsedContent(t *testing.T) {

    jsonBlob := []byte(`{"foobar\u0062": "some value", "Partie du discours\u0062": "Foobar", "Vocable": "BarBaz"}`)

    content := ConvertToParsedContent(jsonBlob)

    expected := datatype.ParsedContent{X: "some value", Discours: "Foobar", Vocabulaire: "BarBaz"}

    TestingCompare(t, content.Discours, expected.Discours)
    TestingCompare(t, content.X, expected.X)
    TestingCompare(t, content.Vocabulaire, expected.Vocabulaire)

}
