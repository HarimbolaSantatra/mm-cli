package datatype

import "fmt"

// Represent the JSON returned by `mm-parsing`
type ParsedContent struct {
    
    // Test
    X string `json:"foobar\u0062"`
    // partie du discours
    Discours string `json:"Partie du discours\u0062"`
    // Explication en malgache
    MlgExplication string `json:"Explications en malgache\u00a0"`
    // Explication en french
    FrExplication string `json:"Explications en fran\u00e7ais\u00a0"`

    Vocabulaire string `json:"Vocable"`
    Morphologie string `json:"Morphologie\u00a0"`
    Analogue string `json:"Analogues\u00a0"`
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
