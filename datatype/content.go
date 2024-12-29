package datatype

// Represent the JSON returned by `mm-parsing`
type ParsedContent struct {
    Discours string // partie du discours
    MlgExplication string // Explication en malgache
    FrExplication string // Explication en french
    Vocabulaire string
    Morphologie string
    Analogue string
}
