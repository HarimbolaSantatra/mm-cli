package datatype

import "fmt"

// Represent the JSON returned by `mm-parsing`
type ParsedContent struct {
    
    Discours string `json:"Part of speech\u00a0"`
    MlgExplication string `json:"Explanations in Malagasy\u00a0"`
    FrExplication string `json:"Explanations in French\u00a0"`
    Vocabulaire string `json:"Vocabulary\u00a0"`
    Morphologie string `json:"Morphology\u00a0"`
    Analogue string `json:"Analogs\u00a0"`
}


// Print debug information (print the value of each field)
func (c ParsedContent) DebugPrint() {
    fmt.Println("Fields of struct ParsedContent:")
    fmt.Printf("- Discours: %s\n", c.Discours)
    fmt.Printf("- MlgExplication: %s\n", c.MlgExplication)
    fmt.Printf("- FrExplication: %s\n", c.FrExplication)
    fmt.Printf("- Vocabulaire: %s\n", c.Vocabulaire)
    fmt.Printf("- Morphologie: %s\n", c.Morphologie)
    fmt.Printf("- Analogue: %s\n", c.Analogue)
}
