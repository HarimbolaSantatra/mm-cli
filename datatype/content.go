package datatype

import "fmt"

// Represent the JSON returned by `mm-parsing`
type ParsedContent struct {
	Speech         string `json:"Part of speech\u00a0"`
	MlgExplication string `json:"Explanations in Malagasy\u00a0"`
	FrExplication  string `json:"Explanations in French\u00a0"`
	Vocabulary     string `json:"Vocabulary\u00a0"`
	Morphology     string `json:"Morphology\u00a0"`
	Analogs        string `json:"Analogs\u00a0"`
	Json           string
}

// Print debug information (print the value of each field)
func (c ParsedContent) DebugPrint() {
	fmt.Println("Fields of struct ParsedContent:")
	fmt.Printf("- Speech: %s\n", c.Speech)
	fmt.Printf("- MlgExplication: %s\n", c.MlgExplication)
	fmt.Printf("- FrExplication: %s\n", c.FrExplication)
	fmt.Printf("- Vocabulary: %s\n", c.Vocabulary)
	fmt.Printf("- Morphology: %s\n", c.Morphology)
	fmt.Printf("- Analogs: %s\n", c.Analogs)
}
