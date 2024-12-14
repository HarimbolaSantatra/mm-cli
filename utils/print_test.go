package utils

import (
    "strings"
    "testing"
)


func TestClean(t *testing.T) {
    const x = "Ito dia texte[3.45] fanao[v.45] test[ okon. ]ma80nde tsara v?]"
    y := Clean(x)
    expected := "Ito dia texte fanao testmande tsara v?]"
    if strings.Compare(y, expected) != 0 {
	t.Errorf("Got: \"%s\"\nExpecting: \"%s\"", y, expected)
    }
}
